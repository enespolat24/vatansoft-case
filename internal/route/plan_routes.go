package routes

import (
	"vatansoft-case/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitPlanRoutes(e *echo.Echo, h *handler.PlanHandler) {
	e.POST("/plans", h.CreatePlan)
	e.GET("/plans", h.GetPlans)
	e.PUT("/plans/:id", h.UpdatePlan)
	e.PUT("/plans/:id/state", h.ChangeState)
	e.GET("/plans/weekly", h.GetWeeklyPlans)
	e.GET("/plans/monthly", h.GetMonthlyPlans)
}
