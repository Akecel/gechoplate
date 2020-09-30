package controller

import (
	"net/http"
	"time"

	e "entity"
	"model"
	"helper"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login verifies the identifiers and connects the user by creating a token. 
func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := model.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	match := CheckPasswordHash(password, res.Password)
	if match != true {
		return c.JSON(http.StatusBadRequest, "Bad password")
	}

	t, err := helper.getJWTToken(res)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Register create a new user in the database
func Register(c echo.Context) error {

}

// Logout delete the token of the current user
func Logout(c echo.Context) error {

}

// CheckPasswordHash compare the hash of the password with the password used by the user
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


