package config

import (
	"github.com/caarlos0/env/v6"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
)

type Local struct {
	KeepLocal    bool   `env:"LOCAL_ONLY" envDefault:"false"`
	Development  bool   `env:"DEVELOPMENT" envDefault:"true"`
	Port         int    `env:"LOCAL_PORT" envDefault:"3000"`
	VaultAddress string `env:"VAULT_ADDRESS" envDefault:"http://vault.vault:8200"`
	RDSAddress   string `env:"RDS_HOSTNAME" envDefault:"postgres.k8sdeploy"`
}

type Config struct {
	Local
	RDS
}

func BuildConfig() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return cfg, bugLog.Errorf("parse env: %+v", err)
	}

	if err := buildDatabase(&cfg); err != nil {
		return cfg, bugLog.Errorf("buildDatabase: %+v", err)
	}

	return cfg, nil
}
