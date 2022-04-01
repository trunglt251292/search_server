package model

import "time"

// CityRaw ...
type CityRaw struct {
	ID           AppID     `bson:"_id" json:"_id"`
	Name         string    `bson:"name" json:"name"`
	Source       string    `bson:"source" json:"source"`
	SearchString string    `bson:"searchString" json:"searchString"`
	Code         int       `bson:"code" json:"code"`
	OsirisID     int       `bson:"osirisId,omitempty" json:"osirisId,omitempty"`
	Slug         string    `bson:"slug" json:"slug"`
	TempID       int       `bson:"tempId" json:"tempId"`
	UpdatedAt    time.Time `bson:"updatedAt" json:"-"`
	Region       string    `bson:"region" json:"region"`
	Checksum     string    `bson:"checkSum"`
	Order        int       `bson:"order"`

	TNCID   int    `bson:"tncId" json:"-"`
	TNCCode string `bson:"tncCode" json:"-"`
}
