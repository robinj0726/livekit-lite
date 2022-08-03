package service

import "livekit-lite/pkg/config"

type LivekitServer struct {
	config *config.Config
}

func NewLivekitServer(conf *config.Config) (s *LivekitServer, err error) {
	s = &LivekitServer{}

	return
}

func (s *LivekitServer) Start() error {
	return nil
}

func (s *LivekitServer) Stop(force bool) error {
	return nil
}
