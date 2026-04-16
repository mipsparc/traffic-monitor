package handler

import "github.com/labstack/echo/v5"

func ReportHandler(c *echo.Context) error {

	return c.JSON(200, "OK")
}
