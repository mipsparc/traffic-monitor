package server

import (
	"github.com/labstack/echo/v5"

	"traffic-monitor/internal/handler"
)

func defineRoutes(e *echo.Echo) {
	e.POST("/api/v1/report", handler.ReportHandler)
}
