package handler

import (
	"log/slog"
	"traffic-monitor/internal/model"
	"traffic-monitor/internal/repository"

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

	errorOccurred := false
	for _, report := range reports.Report {
		err := repository.InsertReport(reports.CameraID, report)
		if err != nil {
			errorOccurred = true
			slog.Error("Failed to insert report: ", report, err)
		}
	}
	if errorOccurred {
		return c.JSON(500, "Server error")
	}

	// add to Valkey to send notification

	return c.JSON(200, reports)
}
