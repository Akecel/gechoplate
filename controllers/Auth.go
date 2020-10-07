package controllers

import (
	"net/http"

	db "gechoplate/database"
	helper "gechoplate/helpers"
	model "gechoplate/models"

	"github.com/labstack/echo/v4"
)

// Login verifies the identifiers and connects the user by creating a token.
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := model.User{}

	if result := db.Gorm.Where("email = ?", email).First(&user); result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}

	match := helper.CheckPasswordHash(password, user.Password)
	if match != true {
		return c.JSON(http.StatusBadRequest, "Bad password")
	}

	t, err := helper.GetJWTToken(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "User connected", map[string]string{
		"token": t,
	}))
}

// Register create a new user in the database
func Register(c echo.Context) error {
	return nil
}

// Logout delete the token of the current user
func Logout(c echo.Context) error {
	return nil
}
