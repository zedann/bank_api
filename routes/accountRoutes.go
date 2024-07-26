package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/zedann/bank_api/handlers"
)

func RegisterAccountRoutes(router *echo.Group) {

	router.GET("/accounts/:id", handlers.HandleGetAccount)
	router.POST("/accounts", handlers.HandleCreateAccount)
	router.DELETE("/accounts/:id", handlers.HandleDeleteAccount)
	router.GET("/accounts", handlers.HandleTransfer)

}
