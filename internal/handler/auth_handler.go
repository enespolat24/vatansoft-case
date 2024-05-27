package handler

import (
	"net/http"
	"time"
	"vatansoft-case/internal/auth"
	"vatansoft-case/internal/model"
	"vatansoft-case/internal/repository"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	UserRepository repository.UserRepository
}

func NewAuthHandler(userRepository repository.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepository: userRepository}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *AuthHandler) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, err := ah.UserRepository.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	if err := auth.VerifyPassword(user.Password, req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password", "password": user.Password, "req.Password": req.Password})
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	expiresAt := time.Now().Add(time.Hour * 24 * 7) // Example: Token expires in 7 days
	if err := auth.(user.ID, token, expiresAt); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (ah *AuthHandler) Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing required fields"})
	}

	exists, _ := ah.UserRepository.UserExistsByEmail(req.Email)
	if exists != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "email already exists"})
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to hash password"})
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := ah.UserRepository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered successfully"})
}
