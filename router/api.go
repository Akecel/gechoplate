package router

import (
	"net/http"

	"gechoplate/controllers"
	"gechoplate/helpers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetAPIRoutes define all apis routes.
func SetAPIRoutes(e *echo.Echo) {

	// Public group
	e.GET("/", controllers.Index)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.SetResponse(http.StatusOK, "All routes", e.Routes()))
	})

	// Authentication routes
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))

	// Refresh token routes
	r.POST("/refresh", controllers.RefreshToken)
}
