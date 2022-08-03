package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port          uint32    `yaml:"port"`
	BindAddresses []string  `yaml:"bind_addresses"`
	RTC           RTCConfig `yaml:"rtc,omitempty"`
}

type RTCConfig struct {
	UDPPort uint32 `yaml:"udp_port,omitempty"`
	TCPPort uint32 `yaml:"tcp_port,omitempty"`
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
