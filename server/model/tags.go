package model

import "time"

type TagRaw struct {
	ID           AppID     `bson:"_id"`
	SearchString string    `bson:"searchString"`
	Type         string    `bson:"type"`
	Name         string    `bson:"name"`
	Color        string    `bson:"color"`
	Active       bool      `bson:"active"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
