package service

import "livekit-lite/pkg/config"

func InitializeServer(conf *config.Config) (*LivekitServer, error) {
	livekitServer, err := NewLivekitServer(conf)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil

}
