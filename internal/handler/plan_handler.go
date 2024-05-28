package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"vatansoft-case/internal/model"
	"vatansoft-case/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type PlanHandler struct {
	PlanRepository repository.PlanRepository
}

func NewPlanHandler(planRepo repository.PlanRepository) *PlanHandler {
	return &PlanHandler{PlanRepository: planRepo}
}

type CreatePlanRequest struct {
	UserID    uint   `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Title     string `json:"title"`
	State     string `json:"state"`
}

type ChangeStateRequest struct {
	State string `json:"state"`
}

type UpdatePlanRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Title     string `json:"title"`
}

func (ph *PlanHandler) CreatePlan(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	userID, err := getUserIdFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	req := new(CreatePlanRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	startTime, err := time.Parse("2006-01-02 15:04", req.StartTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid start time format"})
	}

	endTime, err := time.Parse("2006-01-02 15:04", req.EndTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid end time format"})
	}

	newPlan := &model.Plan{
		UserID:    userID,
		StartTime: startTime,
		EndTime:   endTime,
		Title:     req.Title,
		State:     model.PlanState(req.State),
	}

	overlap, err := ph.PlanRepository.CheckPlanOverlap(newPlan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to check plan overlap"})
	}
	if overlap {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "new plan overlaps with existing plans"})
	}

	if err := ph.PlanRepository.CreatePlan(newPlan); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create plan"})
	}

	return c.JSON(http.StatusCreated, newPlan)
}

func (ph *PlanHandler) GetPlans(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	userID, err := getUserIdFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	plans, err := ph.PlanRepository.GetPlansByUserId(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch plans"})
	}

	return c.JSON(http.StatusOK, plans)
}

func getUserIdFromToken(tokenString string) (uint, error) {
	if tokenString == "" {
		return 0, errors.New("missing Authorization header")
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to parse claims")
	}

	userID := uint(claims["user_id"].(float64))
	return userID, nil
}

func (ph *PlanHandler) ChangeState(c echo.Context) error {
	planID := c.Param("id")

	var req ChangeStateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if !model.IsValidPlanState(string(model.PlanState(req.State))) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid plan state"})
	}

	planIDUint, err := strconv.ParseUint(planID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid plan ID"})
	}

	plan, err := ph.PlanRepository.GetPlanByID(uint(planIDUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "plan not found"})
	}

	plan.State = model.PlanState(req.State)
	if err := ph.PlanRepository.ChangeState(plan.ID, req.State); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update plan"})
	}

	return c.JSON(http.StatusOK, plan)
}

func (ph *PlanHandler) UpdatePlan(c echo.Context) error {
	planID := c.Param("id")

	var req UpdatePlanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	planIDUint, err := strconv.ParseUint(planID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid plan ID"})
	}

	plan, err := ph.PlanRepository.GetPlanByID(uint(planIDUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "plan not found"})
	}

	if req.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04", req.StartTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid start time format"})
		}
		plan.StartTime = startTime
	}

	if req.EndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04", req.EndTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid end time format"})
		}
		plan.EndTime = endTime
	}

	if req.Title != "" {
		plan.Title = req.Title
	}

	if err := ph.PlanRepository.UpdatePlan(plan); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update plan"})
	}

	return c.JSON(http.StatusOK, plan)
}

func (ph *PlanHandler) GetWeeklyPlans(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	userID, err := getUserIdFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	weekStartDate := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	plans, err := ph.PlanRepository.GetWeeklyPlansByUserID(userID, weekStartDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch weekly plans"})
	}
	return c.JSON(http.StatusOK, plans)
}

func (ph *PlanHandler) GetMonthlyPlans(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	userID, err := getUserIdFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	monthStartDate := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	plans, err := ph.PlanRepository.GetMonthlyPlansByUserID(userID, monthStartDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch monthly plans"})
	}

	return c.JSON(http.StatusOK, plans)
}
