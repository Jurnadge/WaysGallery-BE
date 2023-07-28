package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	UserRoutes(e)
	ArtRoutes(e)
	PostRoutes(e)
	HiredRoutes(e)
}
