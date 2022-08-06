package service

import (
	"context"
	"livekit-lite/pkg/config"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/utils"
)

type MyRoomAllocator struct {
	config    *config.Config
	roomStore map[string]*livekit.Room
}

func NewRoomAllocator(conf *config.Config) (RoomAllocator, error) {
	return &MyRoomAllocator{
		config:    conf,
		roomStore: make(map[string]*livekit.Room),
	}, nil
}

func (r *MyRoomAllocator) CreateRoom(ctx context.Context, req *livekit.CreateRoomRequest) (*livekit.Room, error) {
	return &livekit.Room{
		Sid:          utils.NewGuid(utils.RoomPrefix),
		Name:         req.Name,
		CreationTime: time.Now().Unix(),
		TurnPassword: utils.RandomSecret(),
	}, nil

}
