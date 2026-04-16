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

	// ensure 1 or more report exist
	if len(reports.Report) == 0 {
		return c.JSON(400, "Invalid request")
	}

	// ensure all items exist
	if err := c.Validate(reports); err != nil {
		return c.JSON(400, "Invalid request")
	}

	// insert into MySQL

	// check reports are inserted correctly

	// add to Valkey to send notification

	return c.JSON(200, reports)
}
