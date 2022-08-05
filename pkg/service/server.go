package service

import (
	"context"
	"fmt"
	"livekit-lite/pkg/config"
	"net"
	"net/http"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

type LivekitServer struct {
	config *config.Config

	rtcService  *RTCService
	httpServer  *http.Server
	roomManager *RoomManager

	doneChan   chan struct{}
	closedChan chan struct{}
}

func NewLivekitServer(conf *config.Config,
	roomService livekit.RoomService,
	rtcService *RTCService,
	roomManager *RoomManager,
) (s *LivekitServer, err error) {
	s = &LivekitServer{
		config:      conf,
		rtcService:  rtcService,
		roomManager: roomManager,
		closedChan:  make(chan struct{}),
	}

	roomServer := livekit.NewRoomServiceServer(roomService)

	mux := http.NewServeMux()
	mux.Handle(roomServer.PathPrefix(), roomServer)

	s.httpServer = &http.Server{}

	return
}

func (s *LivekitServer) Start() error {
	s.doneChan = make(chan struct{})

	addresses := s.config.BindAddresses
	if addresses == nil {
		addresses = []string{""}
	}

	// ensure we could listen
	listeners := make([]net.Listener, 0)
	for _, addr := range addresses {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, s.config.Port))
		if err != nil {
			return err
		}
		listeners = append(listeners, ln)

	}

	values := []interface{}{
		"portHttp", s.config.Port,
	}
	if s.config.BindAddresses != nil {
		values = append(values, "bindAddresses", s.config.BindAddresses)
	}
	if s.config.RTC.TCPPort != 0 {
		values = append(values, "rtc.portTCP", s.config.RTC.TCPPort)
	}
	if !s.config.RTC.ForceTCP && s.config.RTC.UDPPort != 0 {
		values = append(values, "rtc.portUDP", s.config.RTC.UDPPort)
	} else {
		values = append(values,
			"rtc.portICERange", []uint32{s.config.RTC.ICEPortRangeStart, s.config.RTC.ICEPortRangeEnd},
		)
	}

	logger.Infow("starting LiveKit server", values...)

	for _, ln := range listeners {
		go func(l net.Listener) {
			s.httpServer.Serve(l)
		}(ln)
	}

	// give time for Serve goroutine to start
	time.Sleep(100 * time.Millisecond)

	<-s.doneChan

	// wait for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_ = s.httpServer.Shutdown(ctx)

	// if s.turnServer != nil {
	// 	_ = s.turnServer.Close()
	// }

	// s.roomManager.Stop()
	// s.egressService.Stop()

	close(s.closedChan)
	return nil
}

func (s *LivekitServer) Stop(force bool) {
	close(s.doneChan)

	// wait for fully closed
	<-s.closedChan
}
