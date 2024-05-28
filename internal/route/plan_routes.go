package routes

import (
	"vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitPlanRoutes(e *echo.Echo, h *handler.PlanHandler) {
	e.POST("/plans", h.CreatePlan)
	e.GET("/plans", h.GetPlans)
}
