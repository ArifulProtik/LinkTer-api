package main

import (
	"LinkTer-api/config"
	"LinkTer-api/pkg/logger"
)

func main() {
	_, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.NewApiLogger()
	logger.InitLogger()

}
