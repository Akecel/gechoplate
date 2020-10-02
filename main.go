// API Boilerplate in Go using Echo and Viper
package main

import (
	c "api-goilerplate/config"
	db "api-goilerplate/database"
	r "api-goilerplate/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main launch all part of the project
func main() {

	c.InitConfig()

	db.Connect()

	e := echo.New()

	r.InitRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":80"))
}