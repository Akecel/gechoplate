package middlewares

import (
	"gechoplate/controllers"
	"gechoplate/helpers"

	"net/http"

	"github.com/labstack/echo/v4"
)

// IsAdmin is the middleware function.
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		admin := helpers.RetrieveAdmin(c)
		if admin != true {
			return c.JSON(http.StatusUnauthorized, controllers.SetResponse(http.StatusUnauthorized, "Authorization error", "User is not an administrator"))
		}
		return next(c)
	}
}
