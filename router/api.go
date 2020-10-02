package router

import (
	"gechoplate/controller"
	"gechoplate/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// SetAPIRoutes define all apis routes.
func SetAPIRoutes(e *echo.Echo) {

	// Public group
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, viper.GetString("APP_NAME"), e.Routes()))
	})

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
}