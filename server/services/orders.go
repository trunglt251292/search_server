package services

import (
	"context"
	"encoding/json"
	"fmt"
	"search_server/server/dao"
	"search_server/server/internal/elasticsearch"
	"search_server/server/model"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchOrderAdminByES ...
func (s *Services) SearchOrderAdminByES(ctx context.Context, q *model.CommonQuery) ([]string, int64) {
	var (
		res   = make([]string, 0)
		ids   = make([]string, 0)
		total int64
		//wg    sync.WaitGroup
	)
	qES := q.ConvertESQuery(elasticsearch.OrdersIndex)
	response, err := elasticsearch.Search(*qES)
	if err != nil {
		fmt.Println("Error SearchOrderAdminByES : ", err.Error())
		return res, total
	}

	if err = json.Unmarshal(response.Data, &ids); err != nil {
		fmt.Println(err.Error())
		return res, 0
	}

	if len(ids) == 0 {
		return res, total
	}

	//objIDs := util.ConvertStringsToObjectIDs(ids)
	total = response.Total

	return ids, total
}

func (s *Services) MigrationOrder(isCreate bool) {
	var (
		page, limit, firstPageTotal int64 = 0, 500, 1
		ctx                               = context.Background()
		cond                              = bson.M{}
		total                       int64
	)

	if isCreate {
		_, err := elasticsearch.CreateIndexEs(elasticsearch.OrdersIndex)
		if err != nil {
			panic(err)
		}
	}

	opts := options.Find().SetLimit(firstPageTotal).SetSkip(page * limit)
	docs, _ := dao.OrderFindByCondition(ctx, cond, opts)
	if len(docs) == 0 {
		return
	}
	fmt.Println("*** Starting sync first data order")
	var orderES []*model.OrderES

	for _, item := range docs {
		esDoc := s.GetOrderPayloadES(ctx, &item)
		orderES = append(orderES, esDoc)
	}
	total = total + int64(len(orderES))
	fmt.Println("Total orders: ", total)
	if len(orderES) == 0 {
		return
	}
	_, err := elasticsearch.UpsertDataWithRequest(elasticsearch.OrdersIndex, orderES)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	total = 0
	for {
		opts := options.Find().SetLimit(limit).SetSkip(page * limit)
		docs, _ := dao.OrderFindByCondition(ctx, cond, opts)
		if len(docs) == 0 {
			break
		}
		fmt.Println("*** Starting sync data order page : ", page)
		page++
		var orderES []*model.OrderES

		for _, item := range docs {
			esDoc := s.GetOrderPayloadES(ctx, &item)
			orderES = append(orderES, esDoc)
		}
		total = total + int64(len(orderES))
		fmt.Println("Total orders: ", total)
		if len(orderES) == 0 {
			break
		}
		ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.OrdersIndex, orderES)
		if err != nil {
			fmt.Println("Error : ", err.Error())
		}
		fmt.Println(ok)
	}
	fmt.Println(aurora.Green("Sync order done."))
}

// GetOrderPayloadES ...
func (s *Services) GetOrderPayloadES(ctx context.Context, order *model.OrderRaw) *model.OrderES {
	var (
		userPhone string
		skus      = make([]string, 0)
		tags      = make([]model.TagInfo, 0)
	)
	user, _ := dao.UserFindByID(ctx, order.User)
	if !user.ID.IsZero() && user.Phone != nil {
		userPhone = user.Phone.Full
	}
	items, _ := dao.OrderItemFindAllByOrder(ctx, order.ID)
	for _, item := range items {
		skus = append(skus, item.Sku.Sku)
	}
	if len(order.Tags) > 0 {
		for _, tag := range order.Tags {
			t := dao.GetTagByIDFromCache(ctx, tag)
			if !t.ID.IsZero() {
				tags = append(tags, model.TagInfo{
					ID:   t.ID.Hex(),
					Name: t.Name,
				})
			}
		}
	}
	return order.GetPayloadES(userPhone, skus, tags)
}
