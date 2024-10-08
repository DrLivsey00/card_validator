package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)
			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path

			logger.WithFields(logrus.Fields{
				"method":   method,
				"path":     path,
				"status":   status,
				"duration": duration,
			}).Info("Handled request")
			return err
		}
	}
}
