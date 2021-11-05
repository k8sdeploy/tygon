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

func isPackageEvent(payload interface{}) (bool, *github.PackageEvent) {
	pe := github.PackageEvent{}

	if err := mapstructure.Decode(payload, &pe); err == nil {
		if pe.Package.GetName() != "" {
			return true, &pe
		}
	}

	return false, nil
}

func isPullRequestEvent(payload interface{}) (bool, *github.PullRequestEvent) {
	pe := github.PullRequestEvent{}

	if err := mapstructure.Decode(payload, &pe); err != nil {
		if pe.PullRequest.GetMergeable() {
			return true, &pe
		}
	}

	return false, nil
}

func isReleaseEvent(payload interface{}) (bool, *github.ReleaseEvent) {
	re := github.ReleaseEvent{}

	if err := mapstructure.Decode(payload, &re); err != nil {
		if re.Release.GetTagName() != "" {
			return true, &re
		}
	}

	return false, nil
}
