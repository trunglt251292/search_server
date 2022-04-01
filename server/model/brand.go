package model

import "time"

// BrandRaw ...
type BrandRaw struct {
	ID           AppID     `bson:"_id" json:"_id"`
	Name         string    `bson:"name" json:"name"`
	Slug         string    `bson:"slug" json:"slug"`
	SearchString string    `bson:"searchString" json:"searchString"`
	Active       bool      `bson:"active" json:"active"`
	Desc         string    `bson:"desc" json:"desc"`
	TotalProduct int64     `bson:"totalProduct" json:"totalProduct"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}
