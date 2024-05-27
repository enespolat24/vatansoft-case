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

func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.UserRepository.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, user)
}
