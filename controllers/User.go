package controllers

import (
	"net/http"
	"strconv"

	db "gechoplate/database"
	"gechoplate/helpers"
	"gechoplate/models"

	"github.com/labstack/echo/v4"
)

// GetUser get data of one user
func GetUser(c echo.Context) error {
	user := models.User{}
	userResponse := models.APIUser{}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Parameters error", err.Error()))
	}

	if err := db.Gorm.Model(&user).First(&userResponse, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "User not found", err.Error()))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "User found", userResponse))
}

// GetAllUser get data of all user
func GetAllUser(c echo.Context) error {
	user := models.User{}
	usersResponse := []models.APIUser{}

	if err := db.Gorm.Model(&user).Find(&usersResponse).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Users is empty", err.Error()))
	}

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "Users found", usersResponse))
}

// CreateUser a new user
func CreateUser(c echo.Context) error {
	user := new(models.User)
	userResponse := models.APIUser{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Creation error", err.Error()))
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Creation error", err.Error()))
	}
	user.Password = hashedPassword

	if err := db.Gorm.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Creation error", err.Error()))
	}

	db.Gorm.Model(&user).First(&userResponse, user.ID)

	return c.JSON(http.StatusCreated, SetResponse(http.StatusCreated, "User created", userResponse))
}

// UpdateUser update the user
func UpdateUser(c echo.Context) error {
	user := new(models.User)
	userResponse := models.APIUser{}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Parameters error", err.Error()))
	}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Update error", err.Error()))
	}

	if err := db.Gorm.Select("last_name", "first_name", "email").Where("id = ?", userID).Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Update error", err.Error()))
	}

	db.Gorm.Model(&user).First(&userResponse, user.ID)

	return c.JSON(http.StatusOK, SetResponse(http.StatusOK, "User updated", userResponse))
}

// DeleteUser delelte the user
func DeleteUser(c echo.Context) error {
	user := models.User{}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Parameters error", err.Error()))
	}

	if err := db.Gorm.Delete(&user, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(http.StatusBadRequest, "Delete error", err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}
