package model

import "time"

type (
	// ProductSubCategoriesRaw ...
	ProductSubCategoriesRaw struct {
		ID           AppID     `bson:"_id" json:"_id"`
		Parent       AppID     `bson:"parent" json:"parent"`
		Name         string    `bson:"name" json:"name"`
		SearchString string    `bson:"searchString" json:"-"`
		Order        int       `bson:"order" json:"order"`
		Active       bool      `bson:"active" json:"active"`
		TotalProduct int64     `bson:"totalProduct" json:"totalProduct"`
		CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
	}
)
