package main

import (
	"LinkTer-api/config"
	"fmt"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.Server.Host)

}
