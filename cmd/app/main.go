package main

import (
	"github.com/DrLivsey00/card_checker/internal/config"
	handlers "github.com/DrLivsey00/card_checker/internal/handlers"
	"github.com/DrLivsey00/card_checker/internal/server"
	"github.com/DrLivsey00/card_checker/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	_ = godotenv.Load("configs/.env")
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("/app/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}

	return logger
}

func main() {
	logger := InitLogger()

	if err := godotenv.Load("configs/.env"); err != nil {
		logger.Fatalf("Error loading .env file: %v", err)
	}

	config := config.LoadConfig()
	logger.Infof("Using host port: %s", config.Host)

	services := service.NewService()
	handlers := handlers.NewHandlers(services)
	server := server.NewServer(handlers, logger)
	server.Start(config.Host)
}
