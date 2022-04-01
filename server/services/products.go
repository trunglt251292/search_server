package services

import (
	"context"
	"encoding/json"
	"fmt"
	"search_server/server/config"
	"search_server/server/dao"
	"search_server/server/internal/elasticsearch"
	"search_server/server/internal/util"
	"search_server/server/model"
	"sort"
	"sync"

	"github.com/thoas/go-funk"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchProductAppByES ...
func (s *Services) SearchProductAppByES(ctx context.Context, q *model.CommonQuery) model.ProductAppList {
	var (
		ids = make([]string, 0)
		res = model.ProductAppList{
			Products:    make([]*model.ProductAppResponse, 0),
			Total:       0,
			HasNextPage: false,
		}
	)
	qES := q.ConvertESQuery(elasticsearch.ProductsIndex)
	response, err := elasticsearch.Search(*qES)
	if err != nil {
		fmt.Println("Error SearchProductAppByES : ", err.Error())
		return res
	}

	if err = json.Unmarshal(response.Data, &ids); err != nil {
		fmt.Println("Unmarshal - SearchProductAppByES", err.Error())
		return res
	}

	if len(ids) == 0 {
		return res
	}

	objIDs := util.ConvertStringsToObjectIDs(ids)

	docs, _ := dao.ProductFindByIDs(ctx, objIDs)

	result := convertRawProductToResponseDataWithProductNotShow(ctx, docs, q)

	sort.Slice(result, func(i, j int) bool {
		i1 := funk.IndexOfString(ids, result[i].ID.Hex())
		i2 := funk.IndexOfString(ids, result[j].ID.Hex())
		return i1 < i2
	})

	return model.ProductAppList{
		Products:    result,
		Total:       response.Total,
		HasNextPage: false,
	}
}

// convertRawProductToResponseDataWithProductNotShow ...
func convertRawProductToResponseDataWithProductNotShow(ctx context.Context, list []model.ProductRaw, q *model.CommonQuery) []*model.ProductAppResponse {
	if len(list) == 0 {
		return make([]*model.ProductAppResponse, 0)
	}
	var (
		wg sync.WaitGroup
	)
	res := make([]*model.ProductAppResponse, len(list))
	wg.Add(len(list))
	for index, product := range list {
		go func(p model.ProductRaw, i int) {
			defer wg.Done()
			res[i] = &model.ProductAppResponse{
				ID:               p.ID,
				Name:             p.Name,
				IsOutOfStock:     false,
				PreDiscountPrice: nil,
				Price:            p.Price,
				Info:             p.Info,
				ShareDesc:        p.ShareDesc,
				PendingInactive:  p.PendingInactive,
				InactiveReason:   p.InactiveReason,
				Inventories:      p.Inventories,
				DisplaySKUID:     p.DisplaySKUID,
			}
		}(product, index)
	}
	wg.Wait()
	return res
}

func (s *Services) MigrationProduct() {
	var (
		page, limit int64 = 0, 500
		ctx               = context.Background()
		cond              = bson.M{}
		total       int64
	)
	opts := options.Find().SetLimit(1).SetSkip(page * limit)
	docs, _ := dao.ProductFindByCondition(ctx, cond, opts)
	if len(docs) == 0 {
		return
	}
	fmt.Println("*** Starting sync first data order")
	var productES []*model.ProductESV2

	for _, item := range docs {
		esDoc := s.GetProductPayloadES(ctx, &item)
		productES = append(productES, esDoc)
	}
	total = total + int64(len(productES))
	fmt.Println("Total products: ", total)
	if len(productES) == 0 {
		return
	}
	_, err := elasticsearch.UpsertDataWithRequest(elasticsearch.ProductsIndex, productES)
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	total = 0

	for {
		opts := options.Find().SetLimit(limit).SetSkip(page * limit)
		docs, err := dao.ProductFindByCondition(ctx, cond, opts)
		if err != nil {
			fmt.Println(err.Error())
		}
		if len(docs) == 0 {
			break
		}
		fmt.Println("*** Starting sync data product page : ", page)
		page++
		var productES []*model.ProductESV2

		for _, item := range docs {
			esDoc := s.GetProductPayloadES(ctx, &item)
			productES = append(productES, esDoc)
		}
		total = total + int64(len(productES))
		fmt.Println("Total product: ", total)
		if len(productES) == 0 {
			break
		}
		ok, err := elasticsearch.UpsertDataWithRequest(elasticsearch.ProductsIndex, productES)
		if err != nil {
			fmt.Println("Error : ", err.Error())
		}
		fmt.Println(ok)
	}
	fmt.Println(aurora.Green("Sync product done."))
}

// GetProductPayloadES ...
func (s *Services) GetProductPayloadES(ctx context.Context, product *model.ProductRaw) *model.ProductESV2 {
	esDoc := &model.ProductESV2{
		ID:                product.ID.Hex(),
		Name:              product.Name,
		Categories:        make([]model.ProductCategory, 0),
		Subcategories:     make([]model.ProductCategory, 0),
		Inventories:       make([]model.InventoryCommonInfo, 0),
		Active:            product.Active,
		IsOutOfStock:      product.IsOutOfStock,
		Quantity:          product.Quantity,
		PendingInactive:   product.PendingInactive,
		PendingInactiveAt: product.PendingInactiveAt.Format(config.DateISOFormat),
		InactiveAt:        product.InactiveAt.Format(config.DateISOFormat),
		InactiveReason:    product.InactiveReason,
		CanIssueInvoice:   product.CanIssueInvoice,
		Score:             model.ProductScoreES{},
		Cities:            make([]model.ProductCity, 0),
		CreatedAt:         product.CreatedAt.Format(config.DateISOFormat),
		UpdatedAt:         product.UpdatedAt.Format(config.DateISOFormat),
	}
	if product.Score != nil {
		esDoc.Score = model.ProductScoreES{
			Current: product.Score.Current,
			View:    product.Score.View,
			Order:   int(product.Score.Order),
			Profit:  product.Score.Profit,
		}
	}
	inventory := product.GetInventory()
	if inventory != nil {
		if !inventory.ID.IsZero() {
			esDoc.InventoryId = inventory.ID.Hex()
		}
		esDoc.InventoryName = inventory.Name

		if inventory.Location != nil && inventory.Location.Province != "" {
			city := dao.GetCityBySlugFromCache(ctx, product.Info.Inventory.Location.Province)
			if !city.ID.IsZero() {
				esDoc.Province = city.Code
			}
		}
	}
	// Supplier
	supplier := product.GetSupplier()
	if supplier != nil {
		if !supplier.ID.IsZero() {
			esDoc.SupplierId = supplier.ID.Hex()
		}
		esDoc.SupplierName = supplier.Name
	}

	// Brand
	if !product.Brand.IsZero() {
		esDoc.BrandId = product.Brand.Hex()
		brand := dao.GetBrandByIDFromCache(ctx, product.Brand)
		if !brand.ID.IsZero() {
			esDoc.BrandName = brand.Name
		}
	}
	// Cities
	if len(product.Inventories) > 0 {
		var (
			cities = make([]model.ProductCity, 0)
		)
		for _, inv := range product.Inventories {
			if inv.Location != nil {
				cities = append(cities, model.ProductCity{
					Slug: inv.Location.Province,
				})
			}
		}
		esDoc.Cities = cities
	}
	// Categories
	if len(product.Categories) > 0 {
		for _, cateId := range product.Categories {
			category := dao.GetCategoryByIDFromCache(ctx, cateId)
			if !category.ID.IsZero() {
				esDoc.Categories = append(esDoc.Categories, model.ProductCategory{ID: category.ID.Hex(), Name: category.Name})
			}
		}
	}

	// Sub categories
	if len(product.SubCategories) > 0 {
		for _, subCategoryID := range product.SubCategories {
			subCategories := dao.GetSubCategoryByIDFromCache(ctx, subCategoryID)
			if !subCategories.ID.IsZero() {
				esDoc.Subcategories = append(esDoc.Subcategories, model.ProductCategory{
					ID:   subCategories.ID.Hex(),
					Name: subCategories.Name,
				})
			}
		}
	}

	// Inventories
	if len(product.Inventories) > 0 {
		for _, inv := range product.Inventories {
			esDoc.Inventories = append(esDoc.Inventories, model.InventoryCommonInfo{
				Code: inv.Code,
				ID:   inv.ID.Hex(),
				Name: inv.Name,
			})
		}
	}
	fmt.Println(string(util.ToBytes(esDoc)))
	return esDoc
}
