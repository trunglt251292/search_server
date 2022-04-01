package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"
)

const cityCol = "cities"

// GetCityBySlugFromCache ...
func GetCityBySlugFromCache(ctx context.Context, slug string) model.CityRaw {
	var (
		doc model.CityRaw
	)
	_ = GetCollection(cityCol).FindOne(ctx, bson.M{
		"slug": slug,
	}).Decode(&doc)
	return doc
}
