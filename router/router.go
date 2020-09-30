package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InitRoutes initiates all routes.
func InitRoutes(e *echo.Echo) {
	SetAuthRoutes(e)
	SetAPIRoutes(e)
}