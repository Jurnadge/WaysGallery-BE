package routes

import (
	"juna/handlers"
	"juna/pkg/middleware"
	"juna/pkg/mysql"
	"juna/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	r := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(r)

	e.GET("/profile", middleware.Auth(h.GetUser))
	e.GET("/user", middleware.Auth(h.GetUserDetailByLogin))
	e.GET("/user/:id", middleware.Auth(h.GetUserDetailById))
	e.GET("/users", h.FindUsers)
	e.PATCH("/update-profile", middleware.Auth(middleware.UploadFile(h.UpdateProfile)))
}
