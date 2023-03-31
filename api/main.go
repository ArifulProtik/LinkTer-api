package main

import (
	"LinkTer-api/config"
	"LinkTer-api/internel/controller"
	"LinkTer-api/internel/ent"
	"LinkTer-api/internel/server"
	"LinkTer-api/internel/service"
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
	dbclient := ent.NewDbClient(logger, c)
	services := service.New(logger, dbclient)
	control := controller.New(logger, services)
	group := server.Server.Group("api/v1")
	Router(group, control)

	server.Run(c.Server.Port)
}
