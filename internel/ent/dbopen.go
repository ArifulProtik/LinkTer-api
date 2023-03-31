package ent

import (
	cfg "LinkTer-api/config"
	"LinkTer-api/internel/ent/migrate"
	"LinkTer-api/pkg/logger"
	"context"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDbClient(log logger.Logger, cfg *cfg.Config) *Client {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Postgres.Dbhost, cfg.Postgres.Dbport, cfg.Postgres.DbUser, cfg.Postgres.DBNAME, cfg.Postgres.Dbpassword)
	client, err := Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("DB Connected")
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); !errors.Is(err, nil) {
		log.Fatal(err)
	}
	return client
}
