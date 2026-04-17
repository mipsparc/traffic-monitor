package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func ConsoleHandler(c *echo.Context) error {
	// query param でフィルタリングする
	// filter=camera_id, report_type
	// filterがある場合はvalueで指定
	// sort=severity, time
	// asc=trueならasc, デフォルトはdesc

	return c.JSON(http.StatusOK, "Console handler")
}
