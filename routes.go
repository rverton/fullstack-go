package rvweb

import (
	"net/http"
	"rvweb/components"

	"github.com/labstack/echo/v4"
)

func (hs *HttpServer) routes() {
	hs.Server.GET("/", hs.indexHandler)
}

func (hs *HttpServer) indexHandler(c echo.Context) error {
	return hs.Render(http.StatusOK, c, components.Index())
}
