package middleware

import (
	"net/http"
	"vatansoft-case/internal/auth"
	"vatansoft-case/internal/model"

	"github.com/labstack/echo/v4"
)

func CheckPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*model.User)
			if !hasPermission(user.Role.Name, permission) {
				return echo.NewHTTPError(http.StatusForbidden, "Insufficient permissions")
			}
			return next(c)
		}
	}
}

func hasPermission(roleName string, permission string) bool {
	permissions, ok := auth.RolePermissions[roleName]
	if !ok {
		return false
	}
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}
