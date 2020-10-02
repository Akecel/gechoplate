package router

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes initiates all routes.
func InitRoutes(e *echo.Echo) {
	SetAuthRoutes(e)
	SetAPIRoutes(e)
}
