package server

import (
	"LinkTer-api/pkg/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Server *echo.Echo
	log    logger.Logger
}

func New(log logger.Logger) *Server {
	return &Server{
		Server: echo.New(),
		log:    log,
	}
}

func (s *Server) Run(port string) {
	s.Server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	s.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAccept},
	}))

	go func() {
		s.log.Info(port)
		if err := s.Server.Start(port); err != nil && err != http.ErrServerClosed {
			s.log.Fatal("Shutting Down the Server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		s.log.Fatal(err)
	}

}
