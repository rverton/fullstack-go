package web

import (
	"log"
	"net/http"
	"rvweb/components"

	"github.com/labstack/echo/v4"
)

func (hs *HttpServer) routes() {
	hs.Server.GET("/", hs.indexHandler)
}

func (hs *HttpServer) indexHandler(c echo.Context) error {
	posts, err := hs.Repository.GetPosts()
	if err != nil {
		log.Println("error retrieving posts", err)
	}

	return hs.Render(http.StatusOK, c, components.Index(posts))
}
