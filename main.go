package main

import (
	"airnbn/app/config"
	"airnbn/app/database"
	"airnbn/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)

	database.InitialMigration(dbMysql)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// Route route
	routes.InitRoute(e, dbMysql)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
