package tygon

import (
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/google/go-github/v39/github"
	"github.com/k8sdeploy/tygon/internal/config"
	"github.com/mitchellh/mapstructure"
)

type Tygon struct {
	Config      *config.Config
	EventConfig *EventConfig
}

// const GithubOrgs = "https://api.github.com/orgs/"

// func getOrg(org string) string {
// 	stripedStart := strings.Trim(org, GithubOrgs)
// 	split := strings.Split(stripedStart, "/")
// 	return strings.ToLower(split[0])
// }

func NewTygon(cfg *config.Config) *Tygon {
	return &Tygon{
		Config: cfg,
	}
}

func (t *Tygon) handlePingEvent(p *github.PingEvent) error {
	if err := mapstructure.Decode(p.Hook.Config, &t.EventConfig); err != nil {
		return bugLog.Error(err)
	}

	// org := getOrg(*p.Hook.URL)
	// fmt.Sprintf(org)

	// https://api.github.com/orgs/BugFixes/hooks/326833658

	return nil
}

func (t *Tygon) handlePackageEvent(p *github.PackageEvent) error {
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
