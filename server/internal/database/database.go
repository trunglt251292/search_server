package database

import (
	"search_server/server/config"

	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Connect to mongo server
func Connect(cfg config.MongoCfg) *mongo.Client {
	db, err := mongodb.Connect(cfg.GetConnectOptions())
	if err != nil {
		panic(err)
	}
	return db.Client()
}
