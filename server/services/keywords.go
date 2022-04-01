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

// SearchKeywordAppByES ...
func (s Services) SearchKeywordAppByES(ctx context.Context, q *model.CommonQuery) ([]model.SearchKeywordResponse, int64) {
	var (
		res   = make([]model.SearchKeywordResponse, 0)
		total int64
	)
	qES := q.ConvertESQuery(elasticsearch.KeywordsIndex)
	response, err := elasticsearch.Search(*qES)
	if err != nil {
		fmt.Println("Error SearchKeywordAdminByES : ", err.Error())
		return res, total
	}

	if err = json.Unmarshal(response.Data, &res); err != nil {
		fmt.Println("Error : ", err.Error())
		return res, 0
	}
	total = response.Total
	return res, total
}

func (s *Services) MigrationKeyword() {
	var (
		page, limit    int64 = 0, 500
		firstPageTotal int64 = 1
		ctx                  = context.Background()
		cond                 = bson.M{}
		total          int64
	)
	opts := options.Find().SetLimit(firstPageTotal).SetSkip(0)
	docs, _ := dao.SearchKeywordFindByCondition(ctx, cond, opts)
	if len(docs) == 0 {
		return
	}
	fmt.Println("*** Starting sync data keyword page : ", page)
	page++
	var keywordES []*model.KeywordES

	for _, item := range docs {
		esDoc := s.GetKeywordPayloadES(ctx, &item)
		keywordES = append(keywordES, esDoc)
	}
	total = total + int64(len(keywordES))
	fmt.Println("Total first page keywords: ", total)
	if len(keywordES) == 0 {
		return
	}
	ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.KeywordsIndex, keywordES)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	fmt.Println(aurora.Green(fmt.Sprintf("Sync keywords page %d success : %v", page, ok)))
	total = 0
	for {
		opts := options.Find().SetLimit(limit).SetSkip(page * limit)
		docs, _ := dao.SearchKeywordFindByCondition(ctx, cond, opts)
		if len(docs) == 0 {
			break
		}
		fmt.Println("*** Starting sync data keyword page : ", page)
		page++
		var keywordES []*model.KeywordES

		for _, item := range docs {
			esDoc := s.GetKeywordPayloadES(ctx, &item)
			keywordES = append(keywordES, esDoc)
		}
		total = total + int64(len(keywordES))
		fmt.Println("Total keywords: ", total)
		if len(keywordES) == 0 {
			break
		}
		ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.KeywordsIndex, keywordES)
		if err != nil {
			fmt.Println("Error : ", err.Error())
		}
		fmt.Println(aurora.Green(fmt.Sprintf("Sync keywords page %d success : %v", page, ok)))
	}
	fmt.Println(aurora.Green("Sync keywords done."))
}

// GetKeywordPayloadES ...
func (s *Services) GetKeywordPayloadES(ctx context.Context, keyword *model.SearchKeywordRaw) *model.KeywordES {
	return keyword.GetPayloadES()
}
