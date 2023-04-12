package controller

import (
	"LinkTer-api/config"
	"LinkTer-api/internel/auth"
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
	cfg      *config.Config
	auth     *auth.Auth
}

func New(log logger.Logger, serv *service.Service, cfg *config.Config) *Controller {
	return &Controller{
		log:      log,
		service:  serv,
		validate: lvalidator.New(),
		cfg:      cfg,
		auth:     auth.New(cfg),
	}
}

func (c *Controller) HealthCheck(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"Ping": "Ok",
	})
}

type ErrorResponse struct {
	Msg interface{} `json:"msg"`
}
