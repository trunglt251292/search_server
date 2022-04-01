package services

import (
	"search_server/server/internal/elasticsearch"
	"sync"
)

func (s *Services) MigrationAll() {
	// Delete Multiple indexes
	_, err := elasticsearch.DeleteMultipleIndex(elasticsearch.ListSellyIndex)
	if err != nil {
		panic(err)
	}

	var (
		wg sync.WaitGroup
	)
	wg.Add(len(elasticsearch.ListSellyIndex))
	// Create Index
	for _, index := range elasticsearch.ListSellyIndex {
		go func(i string) {
			defer wg.Done()
			_, err = elasticsearch.CreateIndexEs(i)
			if err != nil {
				panic(err)
			}
		}(index)
	}
	wg.Wait()

	for _, index := range elasticsearch.ListSellyIndex {
		switch index {
		case elasticsearch.ProductsIndex:
			go s.MigrationProduct()
		//case elasticsearch.UsersIndex:
		//	go s.MigrationUser()
		case elasticsearch.OrdersIndex:
			go s.MigrationOrder(false)
		case elasticsearch.KeywordsIndex:
			go s.MigrationKeyword()

		}
	}
}
