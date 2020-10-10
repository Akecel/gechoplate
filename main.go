// API Boilerplate in Go using Echo and Viper
package main

import (
	"gechoplate/config"
	"gechoplate/database"
	"gechoplate/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main launch all part of the project
func main() {

	config.InitConfig()

	database.Connect()
	database.Migrate()
	database.Seed()

	e := echo.New()

	router.InitRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":80"))
}
