package rvweb

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	Server *echo.Echo
}

func (hs *HttpServer) Render(status int, c echo.Context, view templ.Component) error {
	c.Response().Status = status
	return view.Render(context.Background(), c.Response().Writer)
}

func NewHttpServer() *HttpServer {
	hs := &HttpServer{
		Server: echo.New(),
	}

	// setup routes
	hs.routes()

	// assets
	hs.Server.Static("/assets", "assets")

	return hs
}
