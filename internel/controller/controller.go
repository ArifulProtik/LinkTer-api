package controller

import (
	"LinkTer-api/internel/service"
	"LinkTer-api/pkg/logger"
	"LinkTer-api/pkg/lvalidator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	log      logger.Logger
	service  *service.Service
	validate *lvalidator.Lvalidator
}

func New(log logger.Logger, serv *service.Service) *Controller {
	return &Controller{
		log:      log,
		service:  serv,
		validate: lvalidator.New(),
	}
}

func (c *Controller) HealthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"Ping": "Ok",
	})
}