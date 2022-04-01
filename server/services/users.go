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

// SearchUserAdminByES ...
func (s *Services) SearchUserAdminByES(ctx context.Context, q *model.CommonQuery) ([]string, int64) {
	var (
		res   = make([]string, 0)
		ids   = make([]string, 0)
		total int64
		//wg    sync.WaitGroup
	)
	qES := q.ConvertESQuery(elasticsearch.UsersIndex)
	response, err := elasticsearch.Search(*qES)
	if err != nil {
		fmt.Println("Error SearchUserAdminByES : ", err.Error())
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

func (s *Services) MigrationUser() {
	var (
		page, limit int64 = 0, 500
		ctx               = context.Background()
		cond              = bson.M{}
		total       int64
	)
	opts := options.Find().SetLimit(1).SetSkip(page * limit)
	docs, _ := dao.UserFindByCondition(ctx, cond, opts)
	if len(docs) == 0 {
		return
	}
	fmt.Println("*** Starting sync first data user ")
	var userES []*model.UserES

	for _, item := range docs {
		esDoc := s.GetUserPayloadES(ctx, &item)
		userES = append(userES, esDoc)
	}
	total = total + int64(len(userES))
	fmt.Println("Total users: ", total)
	if len(userES) == 0 {
		return
	}
	ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.UsersIndex, userES)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	fmt.Println(aurora.Green(fmt.Sprintf("Sync user page %d success : %v", page, ok)))
	total = 0
	for {
		opts := options.Find().SetLimit(limit).SetSkip(page * limit)
		docs, _ := dao.UserFindByCondition(ctx, cond, opts)
		if len(docs) == 0 {
			break
		}
		fmt.Println("*** Starting sync data user page : ", page)
		page++
		var userES []*model.UserES

		for _, item := range docs {
			esDoc := s.GetUserPayloadES(ctx, &item)
			userES = append(userES, esDoc)
		}
		total = total + int64(len(userES))
		fmt.Println("Total users: ", total)
		if len(userES) == 0 {
			break
		}
		ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.UsersIndex, userES)
		if err != nil {
			fmt.Println("Error : ", err.Error())
		}
		fmt.Println(aurora.Green(fmt.Sprintf("Sync user page %d success : %v", page, ok)))
	}
	fmt.Println(aurora.Green("Sync user done."))
}

// GetUserPayloadES ...
func (s *Services) GetUserPayloadES(ctx context.Context, user *model.UserRaw) *model.UserES {
	var (
		inviter model.AppID
	)
	return user.GetPayloadES(inviter)
}
