package router

import (
	"api-goilerplate/controller"
	
	"github.com/labstack/echo/v4"
)

// SetAuthRoutes define all authentification routes.
func SetAuthRoutes(e *echo.Echo) {
	e.POST("/login", controller.Login)
	e.POST("/regsiter", controller.Register)
	e.GET("/logout", controller.Logout)
}
