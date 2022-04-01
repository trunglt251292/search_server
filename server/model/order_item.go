package model

import "time"

type (
	// OrderItemRaw ...
	OrderItemRaw struct {
		ID           AppID         `bson:"_id" json:"_id"`
		User         AppID         `bson:"user" json:"user"`
		TeamMemberID AppID         `bson:"teamMemberId,omitempty"`
		TeamID       AppID         `bson:"teamId,omitempty"`
		Customer     AppID         `bson:"customer" json:"customer"`
		Order        AppID         `bson:"order" json:"order"`
		Inventory    AppID         `json:"inventory" bson:"inventory"`
		Product      AppID         `bson:"product" json:"product"`
		Date         time.Time     `bson:"date" json:"date"`
		Sku          *OrderItemSku `bson:"sku" json:"sku"`
		Quantity     int64         `bson:"quantity" json:"quantity"`
	}

	// OrderItemSku ...
	OrderItemSku struct {
		ID          AppID       `bson:"_id" json:"_id"`
		Sku         string      `bson:"sku" json:"sku"`
		Type        string      `bson:"type" json:"type"`
		SupplierSku string      `bson:"supplierSku" json:"supplierSku"`
		Name        string      `bson:"name" json:"name"`
		Version     int64       `bson:"version" json:"version"`
		Source      string      `bson:"source" json:"source"`
		Picture     string      `bson:"picture" json:"picture"`
		Properties  []*SKUProps `bson:"properties" json:"properties"`
		UnitCode    string      `bson:"unitCode" json:"unitCode"`
		CanPreorder bool        `bson:"canPreorder" json:"canPreorder"`
		RestockAt   time.Time   `bson:"restockAt,omitempty" json:"restockAt,omitempty"`
	}
)

// SKUProps ...
type SKUProps struct {
	Property AppID `json:"property" bson:"property"`
	Value    AppID `json:"value" bson:"value"`
}
