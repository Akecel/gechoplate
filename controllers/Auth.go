package controllers

import (
	"net/http"

	db "gechoplate/database"
	"gechoplate/helpers"
	"gechoplate/models"

	"github.com/labstack/echo/v4"
)

// Login verifies the identifiers and connects the user by creating a token.
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.User{}

	if err := db.Gorm.Where("email = ?", email).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Connexion error", "User doesn't exist"))
	}

	if match := helpers.CheckPasswordHash(password, user.Password); match != true {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Connexion error", "Bad password"))
	}

	t, rt, err := helpers.GenerateTokenPair(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Connexion error", "JWT error"))
	}

	return c.JSON(http.StatusOK, helpers.SetResponse(http.StatusOK, "User connected", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}

// Register create a new user in the database
func Register(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Register error", err))
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Register error", err))
	}
	user.Password = hashedPassword

	if err := db.Gorm.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "Register error", err))
	}

	return c.JSON(http.StatusCreated, helpers.SetResponse(http.StatusCreated, "User registered", user.ID))
}

// RefreshToken refresh token
func RefreshToken(c echo.Context) error {
	user := models.User{}
	refreshToken := c.FormValue("refresh_token")

	t, rt, err := helpers.RefreshJWTToken(refreshToken, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.SetResponse(http.StatusBadRequest, "JWT Refresh error", "JWT error"))
	}

	return c.JSON(http.StatusOK, helpers.SetResponse(http.StatusOK, "JWT refreshed", map[string]string{
		"refresh_token": rt,
		"token":         t,
	}))
}
