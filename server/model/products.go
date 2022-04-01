package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductRaw ...
type ProductRaw struct {
	ID                  primitive.ObjectID   `bson:"_id" json:"_id"`
	Name                string               `bson:"name" json:"name"`
	SearchString        string               `bson:"searchString" json:"-"`
	Desc                string               `bson:"desc" json:"desc"`
	ShareDesc           string               `bson:"shareDesc" json:"shareDesc"`
	Categories          []primitive.ObjectID `bson:"categories" json:"categories"`
	SubCategories       []primitive.ObjectID `bson:"subCategories" json:"subCategories"`
	Guides              []primitive.ObjectID `bson:"guides" json:"guides"`
	Properties          []primitive.ObjectID `bson:"properties" json:"properties"`
	HighlightProperties []primitive.ObjectID `bson:"highlightProperties" json:"highlightProperties"`
	Info                *ProductInfo         `bson:"info" json:"info"`
	Active              bool                 `bson:"active" json:"active"`
	CreatedAt           time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt           time.Time            `bson:"updatedAt" json:"updatedAt"`
	Price               *SKUPrice            `bson:"price" json:"price"`
	Author              primitive.ObjectID   `bson:"author" json:"author"`
	Score               *ProductScore        `bson:"score" json:"score"`
	Brand               primitive.ObjectID   `bson:"brand,omitempty" json:"brand"`
	IsOutOfStock        bool                 `bson:"isOutOfStock" json:"isOutOfStock"`
	Quantity            int64                `bson:"quantity" json:"quantity"`
	InactiveAt          time.Time            `bson:"inactiveAt,omitempty" json:"inactiveAt,omitempty"`
	PendingInactive     bool                 `bson:"pendingInactive" json:"pendingInactive"`
	PendingInactiveAt   time.Time            `bson:"pendingInactiveAt,omitempty" json:"pendingInactiveAt,omitempty"`
	InactiveReason      string               `bson:"inactiveReason" json:"inactiveReason"`
	CanIssueInvoice     bool                 `bson:"canIssueInvoice"`
	Inventories         []*SKUCommonInfo     `bson:"inventories" json:"inventories"`
	DisplaySKUID        primitive.ObjectID   `bson:"displaySKUId,omitempty" json:"displaySKUId,omitempty"`
}

// GetInventory ...
func (p *ProductRaw) GetInventory() *SKUCommonInfo {
	if p.Info == nil {
		return nil
	}
	return p.Info.Inventory
}

// GetSupplier ...
func (p *ProductRaw) GetSupplier() *SKUCommonInfo {
	if p.Info == nil {
		return nil
	}
	return p.Info.Supplier
}

// ProductInfo ...
type ProductInfo struct {
	Category  *SKUCommonInfo `json:"category" bson:"category"`
	Supplier  *SKUCommonInfo `json:"supplier,omitempty" bson:"supplier,omitempty"`
	Inventory *SKUCommonInfo `json:"inventory" bson:"inventory"`
	Profit    float64        `json:"profit" bson:"profit"`
}

// SKUCommonInfo ...
type SKUCommonInfo struct {
	Code         int                `json:"id,omitempty" bson:"id,omitempty"`
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	SearchString string             `json:"-" bson:"searchString,omitempty"`
	Location     *LocationInventory `json:"location,omitempty" bson:"location,omitempty"`
	MinimumValue float64            `json:"minimumValue,omitempty" bson:"-"`
}

// LocationInventory ...
type LocationInventory struct {
	Address      string         `bson:"address" json:"address"`
	Province     string         `bson:"province" json:"province"`
	ProvinceName string         `bson:"provinceName,omitempty" json:"provinceName,omitempty"`
	District     string         `bson:"district" json:"district"`
	Ward         string         `bson:"ward" json:"ward"`
	FullAddress  string         `bson:"fullAddress,omitempty" json:"fullAddress,omitempty"`
	Location     *MongoLocation `bson:"location" json:"location"`
}

// MongoLocation ...
type MongoLocation struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// ProductScore ...
type ProductScore struct {
	Current float64 `json:"current" bson:"current"`
	View    float64 `json:"view" bson:"view"`
	Order   float64 `json:"order" bson:"order"`
	Profit  float64 `json:"profit" bson:"profit"`
}

// SKUInfo ...
type SKUInfo struct {
	Dimension string         `json:"dimension" bson:"dimension"`
	Weight    float64        `json:"weight" bson:"weight"`
	Inventory *SKUCommonInfo `json:"inventory" bson:"inventory"`
	Category  *SKUCommonInfo `json:"-" bson:"category"`
	Supplier  *SKUCommonInfo `json:"supplier,omitempty" bson:"supplier,omitempty"`
}

// SKUPrice ...
type SKUPrice struct {
	From           float64 `json:"from" bson:"-"`
	Base           float64 `json:"base" bson:"base"`
	Market         float64 `json:"market" bson:"market"`
	Minimum        float64 `json:"minimum" bson:"minimum"`
	Maximum        float64 `json:"maximum" bson:"maximum"`
	Profit         float64 `json:"profit" bson:"profit"`
	Supplier       float64 `json:"supplier" bson:"supplier"`
	WholesaleBonus float64 `json:"wholesaleBonus" bson:"wholesaleBonus"`
}

// ProductAppList ...
type ProductAppList struct {
	Products      []*ProductAppResponse `json:"products"`
	Total         int64                 `json:"total"`
	NextPageToken string                `json:"nextPageToken"`
	HasNextPage   bool                  `json:"-"`
}

// ProductAppResponse ...
type ProductAppResponse struct {
	ID                      AppID            `json:"_id"`
	Name                    string           `json:"name"`
	IsOutOfStock            bool             `json:"isOutOfStock"`
	PreDiscountPrice        *SKUPrice        `json:"preDiscountPrice,omitempty"`
	Price                   *SKUPrice        `json:"price"`
	Info                    *ProductInfo     `json:"info"`
	ShareDesc               string           `json:"shareDesc"`
	ShareImages             []string         `json:"shareImages"`
	PendingInactive         bool             `json:"pendingInactive"`
	InactiveReason          string           `json:"inactiveReason"`
	DoesSupportSellyExpress bool             `json:"doesSupportSellyExpress"`
	Inventories             []*SKUCommonInfo `json:"inventories"`
	DisplaySKUID            AppID            `json:"displaySKUId,omitempty"`
}
