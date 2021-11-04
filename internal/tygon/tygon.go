package tygon

import (
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/google/go-github/v39/github"
	"github.com/k8sdeploy/tygon/internal/config"
	"github.com/mitchellh/mapstructure"
)

type Tygon struct {
	Config    *config.Config
	PingEvent github.Hook
}

func NewTygon(cfg *config.Config) *Tygon {
	return &Tygon{
		Config: cfg,
	}
}

func (t *Tygon) PingEventTriggered(p *github.Hook) error {
	pingConfig := PingEventConfig{}

	if err := mapstructure.Decode(p.Config, &pingConfig); err != nil {
		bugLog.Debugf("failed to decode ping config: %+v", err)
		return err
	}

	bugLog.Infof("Parsed the ping event: %+v", t.PingEvent)

	return nil
}
