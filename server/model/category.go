package model

import "time"

// ProductCategoryRaw ...
type ProductCategoryRaw struct {
	ID           AppID     `bson:"_id" json:"_id"`
	Name         string    `bson:"name" json:"name"`
	Order        int       `bson:"order" json:"order"`
	Active       bool      `bson:"active" json:"active"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
	Featured     bool      `bson:"featured" json:"featured"`
	Color        string    `bson:"color" json:"color"`
	TotalProduct int64     `bson:"totalProduct" json:"totalProduct"`
}
