package routes

import (
	"vatansoft-case/internal/handler"
	"vatansoft-case/internal/middleware"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo, h *handler.UserHandler) {
	e.POST("/users", h.CreateUser, middleware.CheckPermission("create_user"))
	e.GET("/users", h.GetUsers, middleware.CheckPermission("get_users"))
	e.GET("/users/:id", h.GetUser, middleware.CheckPermission("get_user"))
	e.DELETE("/users/:id", h.DeleteUser, middleware.CheckPermission("delete_user"))
}
