package tygon

import (
	"context"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/google/go-github/v39/github"
	"github.com/k8sdeploy/tygon/internal/config"
	"github.com/mitchellh/mapstructure"
)

type Account struct {
	Name   string
	ID     string
	Secret string
}

type Tygon struct {
	Config      *config.Config
	Context     context.Context
	EventConfig *EventConfig
	Account     *Account
}

// const GithubOrgs = "https://api.github.com/orgs/"
//
// //nolint:staticcheck
// func getOrg(org string) string {
// 	stripedStart := strings.Trim(org, GithubOrgs)
// 	split := strings.Split(stripedStart, "/")
// 	return strings.ToLower(split[0])
// }

func NewTygon(cfg *config.Config) *Tygon {
	return &Tygon{
		Config:  cfg,
		Context: context.Background(),
	}
}

func (t *Tygon) handlePingEvent(p *github.PingEvent) error {
	if err := mapstructure.Decode(p.Hook.Config, &t.EventConfig); err != nil {
		return bugLog.Error(err)
	}

	return nil
}

func (t *Tygon) handlePackageEvent(p *github.PackageEvent) error {
	return nil
}

func (t *Tygon) handlePullRequestEvent(p *github.PullRequestEvent) error {
	return nil
}

func (t *Tygon) handleReleaseEvent(p *github.ReleaseEvent) error {
	return nil
}
