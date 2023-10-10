package web

import (
	"net/http"
	"rvweb/components"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (hs *HttpServer) routes() {
	hs.Server.GET("/", hs.indexHandler)
	hs.Server.GET("/post", hs.createHandler)

	hs.Server.POST("/post", hs.insertHandler)
}

func (hs *HttpServer) indexHandler(c echo.Context) error {
	posts, err := hs.Repository.GetPosts()
	if err != nil {
		log.Errorf("failed to get posts: %w", err)
	}

	return hs.Render(http.StatusOK, c, components.Index(posts))
}
func (hs *HttpServer) createHandler(c echo.Context) error {

	// show create form

	return c.String(http.StatusOK, "Hello, World!")
}

func (hs *HttpServer) insertHandler(c echo.Context) error {

	type CreatePost struct {
		Title string `validate:"required"`
	}

	p := new(CreatePost)

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(p); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, e := range validationErrors {
			log.Printf("ERROR: %v: %v", e.Field(), e.Error())
		}

		return c.String(http.StatusOK, "validation failed")
	}

	return c.String(http.StatusOK, "Hello, World!")
}
