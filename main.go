package main

import (
	"fmt"

	"github.com/zahraftrm/mini-project/app/config"
	"github.com/zahraftrm/mini-project/app/database"
	"github.com/zahraftrm/mini-project/app/router"

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

	router.InitRouter(dbMysql, e)

	server := cfg
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", server.SERVER_PORT)))
}
