package dao

import (
	"context"
	"search_server/server/model"
)

func OrderItemFindAllByOrder(ctx context.Context, cond interface{}) ([]model.OrderItemRaw, error) {
	var (
		docs []model.OrderItemRaw
	)
	cursor, err := GetCollection("order-items").Find(ctx, cond)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}
