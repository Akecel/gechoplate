package helper

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

// ParamValidation validate the paramater of the request.
func ParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[0-9]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(http.StatusBadRequest, "param ("+paramKey[k]+") is not a number")
			}
		}

		return next(c)
	}
}

// CheckPasswordHash compare the hash of the password with the password used by the user
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}