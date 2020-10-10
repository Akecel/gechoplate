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
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, controllers.SetResponse(http.StatusOK, "Welcome in "+viper.GetString("APP_NAME"), e.Routes()))
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
