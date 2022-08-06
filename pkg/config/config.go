package config

import (
	"fmt"

	"github.com/livekit/protocol/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port          uint32        `yaml:"port"`
	BindAddresses []string      `yaml:"bind_addresses"`
	RTC           RTCConfig     `yaml:"rtc,omitempty"`
	Room          RoomConfig    `yaml:"room,omitempty"`
	Logging       LoggingConfig `yaml:"logging,omitempty"`

	Development bool `yaml:"development,omitempty"`
}

type RTCConfig struct {
	UDPPort           uint32 `yaml:"udp_port,omitempty"`
	TCPPort           uint32 `yaml:"tcp_port,omitempty"`
	ICEPortRangeStart uint32 `yaml:"port_range_start,omitempty"`
	ICEPortRangeEnd   uint32 `yaml:"port_range_end,omitempty"`

	// for testing, disable UDP
	ForceTCP bool `yaml:"force_tcp,omitempty"`
}

type RoomConfig struct {
	// enable rooms to be automatically created
	AutoCreate         bool        `yaml:"auto_create"`
	EnabledCodecs      []CodecSpec `yaml:"enabled_codecs"`
	MaxParticipants    uint32      `yaml:"max_participants"`
	EmptyTimeout       uint32      `yaml:"empty_timeout"`
	EnableRemoteUnmute bool        `yaml:"enable_remote_unmute"`
	MaxMetadataSize    uint32      `yaml:"max_metadata_size"`
}

type CodecSpec struct {
	Mime     string `yaml:"mime"`
	FmtpLine string `yaml:"fmtp_line"`
}

type LoggingConfig struct {
	logger.Config `yaml:",inline"`
	PionLevel     string `yaml:"pion_level,omitempty"`
}

func NewConfig(confString string) (*Config, error) {
	conf := &Config{}

	if confString != "" {
		if err := yaml.Unmarshal([]byte(confString), conf); err != nil {
			return nil, fmt.Errorf("could not parse config: %v", err)
		}
	}

	return conf, nil
}
