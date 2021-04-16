package contact

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Update(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "PUT /v1/contacts")
}
