package contact

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"signable-sdk/model"
)

func List(ctx echo.Context) error {
	contact := model.Contacts
	return ctx.JSON(http.StatusOK, contact)
}
