package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vatansoft-case/internal/auth"
	"vatansoft-case/internal/middleware"
	"vatansoft-case/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCheckPermission(t *testing.T) {
	testUser := &model.User{
		Role: model.Role{
			Name: "admin",
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", testUser)

	handlerCalled := false
	handler := func(c echo.Context) error {
		handlerCalled = true
		return c.String(http.StatusOK, "Handler called")
	}

	permission := auth.PermissionCreateUser
	middleware.CheckPermission(permission)(handler)(c)
	assert.True(t, handlerCalled, "Handler should be called for permission: "+permission)

	permission = "invalid_permission"
	handlerCalled = false
	middleware.CheckPermission(permission)(handler)(c)
	assert.False(t, handlerCalled, "Handler should not be called for permission: "+permission)
}
