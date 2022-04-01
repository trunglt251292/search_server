package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTagByIDFromCache ...
func GetTagByIDFromCache(ctx context.Context, id model.AppID) model.TagRaw {
	var (
		doc model.TagRaw
	)
	_ = GetCollection("admin-tags").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&doc)
	return doc
}
