package web

import (
	"context"
	"rvweb/repository"

	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	Server *echo.Echo

	Repository *repository.Repository
}
type CustomValidator struct {
	validator *validator.Validate
}

func (hs *HttpServer) Render(status int, c echo.Context, view templ.Component) error {
	c.Response().Status = status
	return view.Render(context.Background(), c.Response().Writer)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func NewHttpServer(db *sqlx.DB) *HttpServer {
	hs := &HttpServer{
		Server:     echo.New(),
		Repository: &repository.Repository{DB: db},
	}

	// setup validation
	hs.Server.Validator = &CustomValidator{validator: validator.New()}

	// setup routes
	hs.routes()

	// assets
	hs.Server.Static("/assets", "assets")

	return hs
}
