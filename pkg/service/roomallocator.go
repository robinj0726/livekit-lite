package service

import (
	"context"
	"livekit-lite/pkg/config"
	"livekit-lite/pkg/routing"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils"
)

type MyRoomAllocator struct {
	config    *config.Config
	router    routing.Router
	roomStore ObjectStore
}

func NewRoomAllocator(conf *config.Config, router routing.Router, rs ObjectStore) (RoomAllocator, error) {
	return &MyRoomAllocator{
		config:    conf,
		router:    router,
		roomStore: rs,
	}, nil
}

func (r *MyRoomAllocator) CreateRoom(ctx context.Context, req *livekit.CreateRoomRequest) (*livekit.Room, error) {
	token, err := r.roomStore.LockRoom(ctx, livekit.RoomName(req.Name), 5*time.Second)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.roomStore.UnlockRoom(ctx, livekit.RoomName(req.Name), token)
	}()

	// find existing room and update it
	rm, err := r.roomStore.LoadRoom(ctx, livekit.RoomName(req.Name))
	if err == ErrRoomNotFound {
		rm = &livekit.Room{
			Sid:          utils.NewGuid(utils.RoomPrefix),
			Name:         req.Name,
			CreationTime: time.Now().Unix(),
			TurnPassword: utils.RandomSecret(),
		}
		applyDefaultRoomConfig(rm, &r.config.Room)
	} else if err != nil {
		return nil, err
	}

	if req.EmptyTimeout > 0 {
		rm.EmptyTimeout = req.EmptyTimeout
	}
	if req.MaxParticipants > 0 {
		rm.MaxParticipants = req.MaxParticipants
	}
	if req.Metadata != "" {
		rm.Metadata = req.Metadata
	}
	if err := r.roomStore.StoreRoom(ctx, rm); err != nil {
		return nil, err
	}

	// check if room already assigned
	_, err = r.router.GetNodeForRoom(ctx, livekit.RoomName(rm.Name))
	if err != routing.ErrNotFound && err != nil {
		return nil, err
	}

	return rm, nil
}

func applyDefaultRoomConfig(room *livekit.Room, conf *config.RoomConfig) {
	room.EmptyTimeout = conf.EmptyTimeout
	room.MaxParticipants = conf.MaxParticipants
	for _, codec := range conf.EnabledCodecs {
		room.EnabledCodecs = append(room.EnabledCodecs, &livekit.Codec{
			Mime:     codec.Mime,
			FmtpLine: codec.FmtpLine,
		})
	}
}
