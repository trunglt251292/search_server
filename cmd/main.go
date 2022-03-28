package main

import (
	"search_server/server/config"
	"search_server/server/initialize"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func init() {
	config.Init()
}

func main() {
	// Echo instance
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Recover())

	initialize.StartServer()

	e.Logger.Fatal(e.Start(config.GetEnv().Port))
}
