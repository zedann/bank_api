package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleAccount(c echo.Context) error {
	return nil
}
func HandleGetAccount(c echo.Context) error {
	id := c.Param("id")

	// db get account

	return c.JSON(http.StatusOK, id)
}
func HandleCreateAccount(c echo.Context) error {
	return nil
}
func HandleDeleteAccount(c echo.Context) error {
	return nil
}
func HandleTransfer(c echo.Context) error {
	return nil
}
