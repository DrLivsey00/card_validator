package server

import (
	"github.com/DrLivsey00/card_checker/internal/handlers"
	midd "github.com/DrLivsey00/card_checker/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	handlers *handlers.Handlers
	server   *echo.Echo
	logger   *logrus.Logger
}

func NewServer(handlers *handlers.Handlers, logger *logrus.Logger) *Server {

	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(midd.Logger(logger))
	logger.SetLevel(logrus.InfoLevel)

	return &Server{
		handlers,
		server,
		logger,
	}
}

func (s *Server) Start(host string) {
	s.InitRoutes()
	s.logger.Fatal(s.server.Start(host))

}

func (s *Server) InitRoutes() {
	g := s.server.Group("/api")
	g.POST("/validate_card", s.handlers.IsValidCard)
}
