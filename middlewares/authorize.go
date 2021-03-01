package middlewares

import (
	"gechoplate/controllers"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// CanManageUser check if the requested user data belongs to the current user or if it is an administrator.
func CanManageUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		
		admin := claims["admin"].(bool)
		id := claims["id"].(uint)
		paramValue := c.ParamValues()
		if admin == true || id == paramValue {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, controllers.SetResponse(http.StatusUnauthorized, "Authorization error", "User is not an administrator and is not authorized to manage this data"))
	}
}

// CanManageData check if the requested user data belongs to the current user or if it is an administrator.
func CanManageData(authorID) (bool) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	admin := claims["admin"].(bool)
	id := claims["id"].(uint)
	if admin == true || id == authorID {
		return true
	}
	return false
}
