package service

import (
	"livekit-lite/pkg/config"
	"livekit-lite/pkg/routing"
)

func InitializeServer(conf *config.Config, currentNode routing.LocalNode) (*LivekitServer, error) {
	router := routing.CreateRouter(currentNode)
	objectStore := createStore()

	roomAllocator, err := NewRoomAllocator(conf, router, objectStore)
	if err != nil {
		return nil, err
	}

	roomService, err := NewRoomService(roomAllocator, objectStore, router)
	if err != nil {
		return nil, err
	}

	roomManager, err := NewLocalRoomManager(conf, objectStore, currentNode, router)
	if err != nil {
		return nil, err
	}

	rtcService := NewRTCService(conf)

	livekitServer, err := NewLivekitServer(conf, roomService, rtcService, roomManager)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil

}

func createStore() ObjectStore {
	return NewLocalStore()
}
