package service

import (
	"LinkTer-api/internel/ent"
	"LinkTer-api/pkg/logger"
)

type Service struct {
	log    logger.Logger
	client *ent.Client
}

func New(log logger.Logger, client *ent.Client) *Service {
	return &Service{
		log:    log,
		client: client,
	}
}
