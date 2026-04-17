package server

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v5"

	"traffic-monitor/internal/handler"
)

func defineRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.POST("/report", handler.ReportHandler)
	//v1.GET("/console", handler.ConsoleHandler)

	// Console
	e.GET("/", handler.IndexHandler)

	// Static content for the console
	assetHandler := http.FileServer(http.FS(os.DirFS("/static")))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
}
