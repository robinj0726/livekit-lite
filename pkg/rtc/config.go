package rtc

import "github.com/pion/webrtc/v3"

type WebRTCConfig struct {
	Configuration webrtc.Configuration
	SettingEngine webrtc.SettingEngine
}
