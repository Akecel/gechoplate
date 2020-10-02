package controllers

import (
	"net/http"

	helper "gechoplate/helpers"
	model "gechoplate/models"

	"github.com/labstack/echo/v4"
)

// Login verifies the identifiers and connects the user by creating a token.
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := model.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	match := helper.CheckPasswordHash(password, res.Password)
	if match != true {
		return c.JSON(http.StatusBadRequest, "Bad password")
	}

	t, err := helper.GetJWTToken(res)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Register create a new user in the database
func Register(c echo.Context) error {
	return nil
}

// Logout delete the token of the current user
func Logout(c echo.Context) error {
	return nil
}
