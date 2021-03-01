package router

import (
	"net/http"

	"gechoplate/controllers"
	"gechoplate/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
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

	r.GET("/user", controllers.GetAllUser, middlewares.IsAdmin)
	r.GET("/user/:id", controllers.GetUser, middlewares.ParamValidation, middlewares.CanManageUser)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser, middlewares.ParamValidation, middlewares.CanManageUser)
	r.DELETE("/user/:id", controllers.DeleteUser, middlewares.ParamValidation, middlewares.CanManageUser)

	// Refresh token routes
	r.POST("/refresh", controllers.RefreshToken)
}
