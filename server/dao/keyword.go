package dao

import (
	"context"
	"search_server/server/model"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchKeywordFindByCondition ...
func SearchKeywordFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.SearchKeywordRaw, err error) {
	cursor, err := GetCollection("search-keywords").Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}
