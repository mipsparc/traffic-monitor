package server

import (
	"github.com/labstack/echo/v5"

	"traffic-monitor/internal/handler"
)

func defineRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.POST("/report", handler.ReportHandler)
	v1.GET("/console", handler.ConsoleHandler)
	// to reduce the number of calls to get full data, just return the latest report_id
	v1.GET("/console/latest", handler.LatestHandler)

	// Static content for the console
	e.File("/", "static/index.html")
}
