package routes

import (
	"vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h *handler.UserHandler) {
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
}
