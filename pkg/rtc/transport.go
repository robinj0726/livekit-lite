package rtc

import (
	"github.com/pion/interceptor"
	"github.com/pion/webrtc/v3"
)

type TransportParams struct {
	Config *WebRTCConfig
}

// PCTransport is a wrapper around PeerConnection, with some helper methods
type PCTransport struct {
	params TransportParams
	pc     *webrtc.PeerConnection
	me     *webrtc.MediaEngine
}

func NewPCTransport(params TransportParams) (*PCTransport, error) {
	t := &PCTransport{
		params: params,
	}

	if err := t.createPeerConnection(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *PCTransport) createPeerConnection() error {
	pc, me, err := newPeerConnection(t.params)
	if err != nil {
		return err
	}

	t.pc = pc
	t.me = me
}

func newPeerConnection(params TransportParams) (*webrtc.PeerConnection, *webrtc.MediaEngine, error) {
	me, err := createMediaEngine()
	if err != nil {
		return nil, nil, err
	}

	se := params.Config.SettingEngine
	ir := &interceptor.Registry{}

	api := webrtc.NewAPI(
		webrtc.WithMediaEngine(me),
		webrtc.WithSettingEngine(se),
		webrtc.WithInterceptorRegistry(ir),
	)
	pc, err := api.NewPeerConnection(params.Config.Configuration)
	return pc, me, err
}
