package routes

import (
	"juna/handlers"
	"juna/pkg/middleware"
	"juna/pkg/mysql"
	"juna/repositories"

	"github.com/labstack/echo/v4"
)

func PostRoutes(e *echo.Group) {
	postRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(postRepository)

	e.GET("/posts", middleware.Auth(h.FindPosts))
	e.GET("/post/:id", middleware.Auth(h.GetPost))
	e.POST("/post", middleware.Auth(middleware.UploadMultipleFile(h.CreatePost)))
}
