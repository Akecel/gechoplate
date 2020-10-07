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
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, "Connexion error", "User doesn't exist"))
	}

	match := helper.CheckPasswordHash(password, user.Password)
	if match != true {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, "Connexion error", "Bad password"))
	}

	t, rt, err := helper.GenerateTokenPair(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, "Connexion error", "JWT error"))
	}

	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "User connected", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}

// Register create a new user in the database
func Register(c echo.Context) error {
	return nil
}

// RefreshToken refresh token
func RefreshToken(c echo.Context) error {
	user := model.User{}
	refreshToken := c.FormValue("refresh_token")

	t, rt, err := helper.RefreshJWTToken(refreshToken, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(http.StatusBadRequest, "JWT Resfresh error", "JWT error"))
	}

	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "JWT refreshed", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}
