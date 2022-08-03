package service

type LivekitServer struct {
}

func NewLivekitServer() (s *LivekitServer, err error) {
	s = &LivekitServer{}

	return
}

func (s *LivekitServer) Start() error {
	return nil
}

func (s *LivekitServer) Stop(force bool) error {
	return nil
}
