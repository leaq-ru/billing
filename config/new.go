package config

import "github.com/kelseyhightower/envconfig"

func New() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	cfg.ServiceName = "billing"
	return
}
