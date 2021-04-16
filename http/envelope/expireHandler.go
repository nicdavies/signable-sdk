package envelope

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Expire(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "GET/v1/envelopes")
}
