package middlewares

import (
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateLoggerOutput() io.Writer {
	// Create a file for logging
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return io.MultiWriter(logFile, os.Stdout)
}

func Logger() echo.MiddlewareFunc {
	// Create a logger instance
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: CreateLoggerOutput(),
	})

	return logger
}
