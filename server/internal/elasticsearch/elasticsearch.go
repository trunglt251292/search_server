package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"search_server/server/config"
	"search_server/server/internal/util"
	"time"

	"github.com/Selly-Modules/natsio"
	"github.com/logrusorgru/aurora"
	"github.com/olivere/elastic"

	es "github.com/Selly-Modules/elasticsearch"
)

var (
	client   *elastic.Client
	ctx      = context.Background()
	esClient *es.Client
)

// InitClientES ...
func InitClientES(apiKey string, nats config.NatsConfig) {
	c, err := es.NewClient(es.Config{
		ApiKey: apiKey,
		Nats: natsio.Config{
			URL:            nats.URL,
			User:           nats.Username,
			Password:       nats.Password,
			RequestTimeout: 3 * time.Minute,
		},
	})
	if err != nil {
		fmt.Println(aurora.Red("CONNECT TO ELASTICSEARCH SERVICE ERR: " + err.Error()))
		panic(err)
	}
	esClient = c
	fmt.Println(aurora.Green("INIT ELASTICSEARCH SERVICE SUCCESS"))
}

func UpsertDataWithPull(indexName string, data interface{}) (ok bool, err error) {
	switch indexName {
	case ProductsIndex:
		return esClient.Pull.ProductUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	case UsersIndex:
		return esClient.Pull.UserUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	case OrdersIndex:
		return esClient.Pull.OrderUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	}
	return false, errors.New("index not support")
}

// DeleteMultipleIndex ...
func DeleteMultipleIndex(indexes []string) (res *es.Response, err error) {
	return esClient.Request.DeleteMultipleIndex(indexes)
}

// CreateIndexEs ...
func CreateIndexEs(indexName string) (res *es.Response, err error) {
	switch indexName {
	case ProductsIndex:
		return esClient.Request.ProductCreateIndex()
	case UsersIndex:
		return esClient.Request.UserCreateIndex()
	case OrdersIndex:
		return esClient.Request.OrderCreateIndex()
	case KeywordsIndex:
		return esClient.Request.KeywordCreateIndex()
	}
	return nil, errors.New("aa index not support")
}

//UpsertDataWithRequest ...
func UpsertDataWithRequest(indexName string, data interface{}) (res *es.Response, err error) {
	switch indexName {
	case ProductsIndex:
		return esClient.Request.ProductUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	case UsersIndex:
		return esClient.Request.UserUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	case OrdersIndex:
		return esClient.Request.OrderUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	case KeywordsIndex:
		return esClient.Request.KeywordUpsert(es.Payload{
			Index: indexName,
			Data:  util.ToBytes(data),
		})
	}
	return nil, errors.New("index not support")
}

func Search(esQuery es.ESQuery) (res *es.Response, err error) {
	switch esQuery.Index {
	case ProductsIndex:
		return esClient.Request.ProductSearch(esQuery)
	case UsersIndex:
		return esClient.Request.UserSearch(esQuery)
	case OrdersIndex:
		return esClient.Request.OrderSearch(esQuery)
	case KeywordsIndex:
		return esClient.Request.KeywordSearch(esQuery)
	}
	return nil, errors.New("index not support")
}

// Init ...
func Init(url, username, pwd string) {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(username, pwd), elastic.SetSniff(false))
	if err != nil {
		fmt.Println(aurora.Red("CONNECT TO ELASTICSEARCH ERR: " + err.Error()))
		panic(err)
	}
	info, code, err := client.Ping(url).Do(ctx)
	if err != nil {
		fmt.Println(aurora.Red("PING ELASTICSEARCH ERR: " + err.Error()))
		panic(err)
	}
	fmt.Println(aurora.Green("CONNECTED TO ELASTICSEARCH: " + url))
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := client.ElasticsearchVersion(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion)
}

// GetClient ...
func GetClient() *elastic.Client {
	return client
}
