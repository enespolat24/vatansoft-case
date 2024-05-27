package routes

import (
	handlers "vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/users", handlers.CreateUser)
	// e.GET("/users", handlers.GetUsers)
	// e.GET("/users/:id", handlers.GetUser)
	// e.PUT("/users/:id", handlers.UpdateUser)
	// e.DELETE("/users/:id", handlers.DeleteUser)
}
