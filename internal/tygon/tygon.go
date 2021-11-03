package tygon

import (
	"github.com/k8sdeploy/tygon/internal/config"
)

type Tygon struct {
	Config *config.Config
}

func NewTygon(cfg *config.Config) *Tygon {
	return &Tygon{
		Config: cfg,
	}
}
