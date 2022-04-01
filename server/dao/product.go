package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProductFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.ProductRaw, err error) {
	cursor, err := GetCollection("products").Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

func ProductFindByIDs(ctx context.Context, ids []model.AppID) ([]model.ProductRaw, error) {
	cond := bson.M{"_id": bson.M{"$in": ids}}
	return ProductFindByCondition(ctx, cond)
}
