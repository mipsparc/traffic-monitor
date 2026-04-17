package handler

import (
	"net/http"
	"traffic-monitor/internal/repository"

	"github.com/labstack/echo/v5"
)

func LatestHandler(c *echo.Context) error {
	reportID, err := repository.LatestReportID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(200, reportID)
}
