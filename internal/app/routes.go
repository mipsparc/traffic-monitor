package server

import (
	"github.com/labstack/echo/v5"

	"traffic-monitor/internal/handler"
)

func defineRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.POST("/report", handler.ReportHandler)
}
