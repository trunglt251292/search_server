package model

// ProductESV2 ...
type ProductESV2 struct {
	ID                string                `json:"id"`
	Name              string                `json:"name"`
	Categories        []ProductCategory     `json:"categories"`
	Subcategories     []ProductCategory     `json:"subCategories"`
	Province          int                   `json:"province"`
	InventoryId       string                `json:"inventoryId"`
	InventoryName     string                `json:"inventoryName"`
	Active            bool                  `json:"active"`
	SupplierId        string                `json:"supplierId"`
	SupplierName      string                `json:"supplierName"`
	BrandId           string                `json:"brandId"`
	BrandName         string                `json:"brandName"`
	Score             ProductScoreES        `json:"score"`
	Cities            []ProductCity         `json:"cities"`
	IsOutOfStock      bool                  `json:"isOutOfStock"`
	Quantity          int64                 `json:"quantity"`
	InactiveAt        string                `json:"inactiveAt"`
	PendingInactive   bool                  `json:"pendingInactive"`
	PendingInactiveAt string                `json:"pendingInactiveAt"`
	InactiveReason    string                `json:"inactiveReason"`
	CanIssueInvoice   bool                  `json:"canIssueInvoice"`
	CreatedAt         string                `json:"createdAt"`
	UpdatedAt         string                `json:"updatedAt"`
	Inventories       []InventoryCommonInfo `json:"inventories"`
}

// ProductCategory ...
type ProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InventoryCommonInfo ...
type InventoryCommonInfo struct {
	Code int    `json:"code" bson:"code"`
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

// ProductCity ...
type ProductCity struct {
	Slug string `json:"slug"`
}

// ProductScoreES ...
type ProductScoreES struct {
	Current float64 `json:"current"`
	View    float64 `json:"view"`
	Order   int     `json:"order"`
	Profit  float64 `json:"profit"`
}
