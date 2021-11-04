package tygon

import (
	"github.com/google/go-github/v39/github"
	"github.com/mitchellh/mapstructure"
)

type EventConfig struct {
	ContentType string `mapstructure:"content_type"`
	InsecureSSL string `mapstructure:"insecure_ssl"`
	Scret       string `mapstructure:"secret"`
	URL         string `mapstructure:"url"`
}

func isPingEvent(payload interface{}) (bool, *github.PingEvent) {
	pe := github.PingEvent{}

	if err := mapstructure.Decode(payload, &pe); err == nil {
		if pe.GetZen() != "" {
			return true, &pe
		}
	}

	return false, nil
}
