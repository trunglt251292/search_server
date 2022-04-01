package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCategoryByIDFromCache(ctx context.Context, id model.AppID) model.ProductCategoryRaw {
	// Get value from cache
	var (
		doc model.ProductCategoryRaw
	)
	_ = GetCollection("product-categories").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&doc)
	return doc
}
