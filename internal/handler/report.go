package handler

import (
	"log/slog"
	"net/http"
	"traffic-monitor/internal/model"
	"traffic-monitor/internal/repository"

	"github.com/labstack/echo/v5"
)

func ReportHandler(c *echo.Context) error {
	var reports model.Reports
	if err := c.Bind(&reports); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// ensure 1 or more report exist
	if len(reports.Report) == 0 {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// ensure all items exist
	if err := c.Validate(reports); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// if the report with the same UUID sent again, ignored by database and logged as error
	errorOccurred := false
	for _, report := range reports.Report {
		if !report.ReportType.Validate() {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		err := repository.InsertReport(reports.CameraID, report)
		if err != nil {
			errorOccurred = true
			slog.Error("Failed to insert report: ", report, err)
		}
	}
	if errorOccurred {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	return c.JSON(http.StatusOK, "OK")
}
