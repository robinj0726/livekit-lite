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
