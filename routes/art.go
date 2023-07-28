package routes

import (
	"juna/handlers"
	"juna/pkg/middleware"
	"juna/pkg/mysql"
	"juna/repositories"

	"github.com/labstack/echo/v4"
)

func ArtRoutes(e *echo.Group) {
	artRepository := repositories.RepositoryArt(mysql.DB)
	h := handlers.HandlerArt(artRepository)

	e.GET("/art/:id", middleware.Auth(h.GetArtByUserId))
	e.POST("/create-art", middleware.Auth(middleware.UploadFile(h.CreateArt)))
}
