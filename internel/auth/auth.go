package auth

import "LinkTer-api/config"

type Auth struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Auth {
	return &Auth{
		cfg: cfg,
	}
}
