package handler

import (
	"traffic-monitor/internal/model"

	"github.com/labstack/echo/v5"
)

func ReportHandler(c *echo.Context) error {
	var reports model.Reports
	if err := c.Bind(&reports); err != nil {
		return c.JSON(400, "Invalid request")
	}

	// validate first

	// validate on each report

	// insert into MySQL

	// check reports are inserted correctly

	// add to Valkey to send notification

	return c.JSON(200, reports)
}
