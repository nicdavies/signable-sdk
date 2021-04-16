package branding

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Update(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "POST /v1/users")
}

func UpdateEmails(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "POST /v1/users")
}
