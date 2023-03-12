package config

import (
	"fmt"

	"github.com/caarlos0/env/v7"
)

// Config Holds all Configuration Details ex: DB Details, App details, Various keys
type Config struct {
	App     string `env:"APP"`
	version string `env:"VERSION"`
	Author  string `env:"AUTHOR"`
	Server
	Postgres
}

type Server struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
}

type Postgres struct {
	Dbhost     string `env:"DBHOST"`
	Dbport     string `env:"DBPORT"`
	Dbpassword string `env:"DBPASS"`
	DbUser     string `env:"DBUSER"`
	DBNAME     string `env:"DBNAME`
}

// New Return New config for the application.
func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	return &cfg, nil
}
