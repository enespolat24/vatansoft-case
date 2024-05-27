package handler

import (
	"net/http"
	"vatansoft-case/internal/model"
	"vatansoft-case/internal/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository: userRepository}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username, email, and password are required"})
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.UserRepository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.UserRepository.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
