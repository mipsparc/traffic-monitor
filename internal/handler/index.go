package handler

import "github.com/labstack/echo/v5"

func IndexHandler(c *echo.Context) error {
	return c.String(200, "Hello, World!")
}
