package routes

import (
	"vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo, h *handler.UserHandler) {
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUser)
	e.DELETE("/users/:id", h.DeleteUser)
}
