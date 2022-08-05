package service

import "livekit-lite/pkg/config"

func InitializeServer(conf *config.Config) (*LivekitServer, error) {

	roomService, err := NewRoomService()
	if err != nil {
		return nil, err
	}

	roomManager, err := NewLocalRoomManager(conf)
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
