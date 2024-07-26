package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zedann/bank_api/routes"
)

type APIServer struct {
	listenAddr string
}

func NewServer(listenAddr string) *APIServer {

	return &APIServer{
		listenAddr: listenAddr,
	}

}

func (s *APIServer) Run() {

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	v1 := e.Group("/api/v1")

	routes.RegisterAccountRoutes(v1)

	e.Logger.Fatal(e.Start(s.listenAddr))
}
