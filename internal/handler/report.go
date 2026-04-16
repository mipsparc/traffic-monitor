package handler

import (
	"github.com/labstack/echo/v5"
)

func ReportHandler(c *echo.Context) error {
	// JSONのReportsを受け取ってunmarshalする
	//reports := &model.Reports{}

	// validate first

	// validate on each report

	// insert into MySQL

	// check reports are inserted correctly

	// add to Valkey to send notification

	return c.JSON(200, "OK")
}
