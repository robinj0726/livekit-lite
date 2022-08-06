package service

import (
	"context"
	"livekit-lite/pkg/config"
	"livekit-lite/pkg/routing"
	"livekit-lite/pkg/rtc"
	"sync"

	"github.com/livekit/protocol/livekit"
)

type RoomManager struct {
	lock sync.RWMutex

	config      *config.Config
	rtcConfig   *rtc.WebRTCConfig
	currentNode routing.LocalNode
	router      routing.Router
	roomStore   ObjectStore

	rooms map[livekit.RoomName]*rtc.Room
}

func NewLocalRoomManager(
	conf *config.Config,
	roomStore ObjectStore,
	currentNode routing.LocalNode,
	router routing.Router,
) (*RoomManager, error) {
	rtcConf, err := rtc.NewWebRTCConfig(conf)
	if err != nil {
		return nil, err
	}

	r := &RoomManager{
		config:      conf,
		rtcConfig:   rtcConf,
		currentNode: currentNode,
		router:      router,
		roomStore:   roomStore,

		rooms: make(map[livekit.RoomName]*rtc.Room),
	}

	return r, nil
}

func (r *RoomManager) GetRoom(_ context.Context, roomName livekit.RoomName) *rtc.Room {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.rooms[roomName]
}
