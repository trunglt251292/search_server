package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"
)

// GetBrandByIDFromCache ...
func GetBrandByIDFromCache(ctx context.Context, id model.AppID) model.BrandRaw {
	var (
		doc model.BrandRaw
	)
	_ = GetCollection("brands").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&doc)
	return doc
}
