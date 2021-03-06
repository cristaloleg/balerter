package config

import (
	"fmt"
	"strings"
)

type ChannelDiscord struct {
	Name      string `json:"name" yaml:"name"`
	Token     string `json:"token" yaml:"token"`
	ChannelID int64  `json:"channelId" yaml:"channelId"`
}

func (cfg *ChannelDiscord) Validate() error {
	if strings.TrimSpace(cfg.Name) == "" {
		return fmt.Errorf("name must be not empty")
	}

	if strings.TrimSpace(cfg.Token) == "" {
		return fmt.Errorf("token must be not empty")
	}

	if cfg.ChannelID < 1 {
		return fmt.Errorf("channel id must be not empty")
	}

	return nil
}
