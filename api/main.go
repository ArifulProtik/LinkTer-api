package main

import (
	"LinkTer-api/config"
	"LinkTer-api/internel/server"
	"LinkTer-api/pkg/logger"
)

func main() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.NewApiLogger()
	logger.InitLogger()
	logger.Info(c.Server.Port)
	server := server.New(logger)
	server.Run(c.Server.Port)
}
