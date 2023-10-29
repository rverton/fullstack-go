package web

import (
	"log/slog"
	"net/http"
	"rvweb/components"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (hs *HttpServer) routes() {
	hs.Server.GET("/", hs.indexHandler)
	hs.Server.GET("/post", hs.createHandler)

	hs.Server.POST("/post", hs.insertHandler)
}

func (hs *HttpServer) indexHandler(c echo.Context) error {
	posts, err := hs.Repository.GetPosts()
	if err != nil {
		slog.Error("failed to get posts: %w", err)
	}

	return hs.Render(http.StatusOK, c, components.Index(posts))
}
func (hs *HttpServer) createHandler(c echo.Context) error {

	return hs.Render(http.StatusOK, c, components.CreatePost(nil, nil))
}

func (hs *HttpServer) insertHandler(c echo.Context) error {

	type CreatePost struct {
		Title string `validate:"required" form:"title"`
		Body  string `validate:"required" form:"body"`
	}

	p := new(CreatePost)

	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(p); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := convertErrors(validationErrors)

		prefill := map[string]string{
			"Title": p.Title,
			"Body":  p.Body,
		}

		return hs.Render(http.StatusBadRequest, c, components.CreatePost(prefill, errors))
	}

	if _, err := hs.Repository.CreatePost(p.Title, p.Body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/")
}
