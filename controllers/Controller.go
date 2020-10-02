package controllers

import (
	"net/http"

	helper "gechoplate/helpers"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Index display all routes of the application.
func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, helper.SetResponse(http.StatusOK, "Welcome in "+viper.GetString("APP_NAME"), "Use "+viper.GetString("APP_URL")+"/routes to see all routes of the application"))
}
