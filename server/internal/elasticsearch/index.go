package elasticsearch

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/logrusorgru/aurora"
	"github.com/olivere/elastic"
)

// Index name
const (
	ViAnalyzer    = "vi_analyzer"
	ProductsIndex = "selly_products"
	UsersIndex    = "selly_users"
	OrdersIndex   = "selly_orders"
	KeywordsIndex = "selly_keywords"
)

var (
	ListSellyIndex = []string{
		ProductsIndex,
		UsersIndex,
		OrdersIndex,
		KeywordsIndex,
	}
)

type AnalyzerBody struct {
	Analyzer string `json:"analyzer"`
	Text     string `json:"text"`
}

// Indexer ...
type Indexer interface {
	GetID() string
}

func indexIsExists(indexName string) bool {
	ctx := context.Background()
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	return exists
}

func GetViTokens(text string) []elastic.AnalyzeToken {
	body := AnalyzerBody{
		Analyzer: ViAnalyzer,
		Text:     text,
	}
	res, err := client.IndexAnalyze().Analyzer(ViAnalyzer).BodyJson(body).Do(context.Background())
	if err != nil {
		fmt.Println("Error :", err)
		return make([]elastic.AnalyzeToken, 0)
	}
	return res.Tokens
}

// CreateIndex ...
func CreateIndex(indexName, mappingFileName string) {
	if indexIsExists(indexName) {
		fmt.Printf("Index %s is existed, skip create.\n", indexName)
		return
	}
	ctx := context.Background()
	m, err := ioutil.ReadFile("internal/module/elasticsearch/mapping/" + mappingFileName)
	if err != nil {
		panic(err)
	}
	_, err = client.CreateIndex(indexName).BodyString(string(m)).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create index %s success\n", indexName)
}

// BulkIndex bulk index documents
func BulkIndex(name string, docs []Indexer) error {
	br := client.Bulk()

	for _, doc := range docs {
		bi := elastic.NewBulkIndexRequest().Index(name).Id(doc.GetID()).Doc(doc)
		br.Add(bi)
	}

	_, err := br.Do(context.Background())
	if err != nil {
		fmt.Println(aurora.Red(fmt.Sprintf("Bulk index %s err: %v", name, err)))
		return err
	}

	return nil
}
