package handlers

import (
	"net/http"
	"vatansoft-case/internal/database"
	"vatansoft-case/internal/model"

	"github.com/labstack/echo/v4"
)

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, user)
}
