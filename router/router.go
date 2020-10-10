package router

import (
	"gechoplate/helpers"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

// InitRoutes initiates all routes.
func InitRoutes(e *echo.Echo) {
	SetAPIRoutes(e)
}

// ParamValidation validate the parameter of the request.
func ParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[0-9]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Validation Error", "Param ("+paramKey[k]+") is not a number"))
			}
		}

		return next(c)
	}
}
