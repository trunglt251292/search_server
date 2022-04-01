package dao

import "go.mongodb.org/mongo-driver/mongo"

var (
	db *mongo.Database
)

func SetMongoDB(d *mongo.Database) {
	db = d
}

func GetCollection(col string) *mongo.Collection {
	return db.Collection(col)
}
