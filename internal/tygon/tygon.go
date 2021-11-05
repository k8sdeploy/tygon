package tygon

import (
	"context"
	"strings"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/google/go-github/v39/github"
	"github.com/k8sdeploy/tygon/internal/config"
	"github.com/mitchellh/mapstructure"
)

type Account struct {
	Name string
	ID   int
}

type Tygon struct {
	Config      *config.Config
	Context     context.Context
	EventConfig *EventConfig
	Account     *Account
}

const GithubOrgs = "https://api.github.com/orgs/"

//nolint:staticcheck
func getOrg(org string) string {
	stripedStart := strings.Trim(org, GithubOrgs)
	split := strings.Split(stripedStart, "/")
	return strings.ToLower(split[0])
}

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

	org := getOrg(*p.Hook.URL)
	if err := t.associateAccountWithOrganization(org); err != nil {
		return bugLog.Errorf("failed to associate the ping: %+v", err)
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

// func (t *Tygon) validateSecret(secretHash string, payload []byte) (bool, error) {
// 	gHash := hmac.New(sha256.New, []byte("abaf8d42-a9b0-401a-a7bb-f32a074f9e3d"))
// 	gHash.Write(payload)
// 	sum := hex.EncodeToString(gHash.Sum(nil))
// 	if hmac.Equal([]byte(fmt.Sprintf("sha256=%s", sum)), []byte(secretHash)) {
// 		fmt.Print("\nthey are the same\n")
// 	} else {
// 		fmt.Print("\nthey are different\n")
// 		fmt.Printf("\tsum: %s\n, \tkey: %s\n, \tprekey: %s\n", sum, secretHash, payload)
// 	}
//
// 	return true, nil
// }
