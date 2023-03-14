package main

import (
	"LinkTer-api/config"
	"LinkTer-api/pkg/logger"
	"LinkTer-api/pkg/lvalidator"
	"fmt"
)

func main() {
	_, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.NewApiLogger()
	logger.InitLogger()
	vv := lvalidator.New()
	type User struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
	u := User{
		Email: "Hello WOrld",
	}
	data := vv.Struct(u)
	if data != nil {
		fmt.Println(data)
	}
}
