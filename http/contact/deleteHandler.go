package contact

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Delete(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "DELETE /v1/contacts")
}
