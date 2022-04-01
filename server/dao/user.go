package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserFindByID(ctx context.Context, id primitive.ObjectID) (doc model.UserRaw, err error) {
	err = GetCollection("users").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&doc)
	return doc, err
}

func UserFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.UserRaw, err error) {
	cursor, err := GetCollection("users").Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}
