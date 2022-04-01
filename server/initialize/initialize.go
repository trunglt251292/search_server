package initialize

import (
	"fmt"
	"search_server/server/config"
	"search_server/server/dao"
	"search_server/server/internal/database"
	"search_server/server/internal/elasticsearch"
	"search_server/server/route"

	"github.com/logrusorgru/aurora"

	"github.com/labstack/echo"
)

// InitDB ...
func InitDB() {
	dbMain := config.GetEnv().Mongo
	db := database.Connect(dbMain)
	fmt.Println(aurora.Green("*** Database name: " + dbMain.DBName))

	dao.SetMongoDB(db.Database(dbMain.DBName))
}

// StartServer ...
func StartServer(e *echo.Echo) {
	InitDB()
	cfg := config.GetEnv()
	esCfg := cfg.Elasticsearch
	elasticsearch.Init(esCfg.URL, esCfg.Username, esCfg.Password)
	elasticsearch.InitClientES(config.GetEnv().ApiKey, cfg.Nats)

	route.Init(e)
}
