package logger

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Fatal(err error) {
	f, _ := os.OpenFile("/var/log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log := echo.New().Logger
	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.Fatal(err)
}

func LoggerWithConfig() echo.MiddlewareFunc {
	f, _ := os.OpenFile("/var/log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(io.MultiWriter(os.Stdout, f, os.Stderr))
	defer f.Close()
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: log.Output(),
	})
}
