package team

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "POST /v1/users")
}
