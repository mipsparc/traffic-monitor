package handler

import (
	"net/http"
	"slices"
	"traffic-monitor/internal/repository"

	"github.com/labstack/echo/v5"
)

func ConsoleHandler(c *echo.Context) error {
	// filter by camera_id or report_type
	filter := c.QueryParam("filter")
	filterValue := ""
	if filter != "" {
		if !slices.Contains([]string{"camera_id", "report_type"}, filter) {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		filterValue = c.QueryParam("filter_value")
		if filterValue == "" {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
	}

	// sort by severity or time
	sort := c.QueryParam("sort")
	if sort != "" {
		if !slices.Contains([]string{"severity", "time"}, sort) {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
	} else {
		sort = "time"
	}

	// when sorted, asc or desc
	asc := c.QueryParam("asc")
	ascDesc := "DESC"
	if asc == "true" {
		ascDesc = "ASC"
	}

	result, err := repository.SearchReport(filter, filterValue, sort, ascDesc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Server error")
	}

	if len(*result) == 0 {
		return c.JSON(http.StatusOK, "")
	}

	return c.JSON(http.StatusOK, result)
}
