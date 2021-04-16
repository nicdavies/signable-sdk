package template

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func List(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "GET/v1/templates")
}
