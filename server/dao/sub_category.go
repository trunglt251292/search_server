package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSubCategoryByIDFromCache(ctx context.Context, id model.AppID) model.ProductSubCategoriesRaw {
	// Get value from cache
	var (
		doc model.ProductSubCategoriesRaw
	)
	_ = GetCollection("product-sub-categories").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&doc)
	return doc
}
