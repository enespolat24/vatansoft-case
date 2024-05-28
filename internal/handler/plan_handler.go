package handler

import (
	"errors"
	"net/http"
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
