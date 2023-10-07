package web

import (
	"context"
	"rvweb/repository"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	Server *echo.Echo

	Repository *repository.Repository
}

func (hs *HttpServer) Render(status int, c echo.Context, view templ.Component) error {
	c.Response().Status = status
	return view.Render(context.Background(), c.Response().Writer)
}

func NewHttpServer(db *sqlx.DB) *HttpServer {
	hs := &HttpServer{
		Server:     echo.New(),
		Repository: &repository.Repository{DB: db},
	}

	// setup routes
	hs.routes()

	// assets
	hs.Server.Static("/assets", "assets")

	return hs
}
