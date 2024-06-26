package routes

import (
	"vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitAuthRoutes(e *echo.Echo, h *handler.AuthHandler) {
	e.POST("/auth/login", h.Login)
	e.POST("/auth/register", h.Register)
	e.GET("/auth/me", h.Me)
	e.PUT("/auth/update-me", h.UpdateMyInformations)
}
