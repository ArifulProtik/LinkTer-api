package main

import (
	"LinkTer-api/internel/controller"

	"github.com/labstack/echo/v4"
)

func Router(group *echo.Group, control *controller.Controller) {
	group.GET("/", control.HealthCheck)
	group.POST("/signup", control.Signup)
	group.POST("/signin", control.SignIn)

}
