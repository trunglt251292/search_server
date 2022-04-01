package model

import (
	"search_server/server/config"
	"time"
)

type (
	// SearchKeywordRaw ...
	SearchKeywordRaw struct {
		ID         AppID     `bson:"_id"`
		Vietnamese string    `bson:"vietnamese"`
		Keyword    string    `bson:"keyword"`
		TimeUnit   TimeUnit  `bson:"timeUnit"`
		Score      int64     `bson:"score"`
		Source     string    `bson:"source"`
		CreatedAt  time.Time `bson:"createdAt"`
		UpdatedAt  time.Time `bson:"updatedAt"`
	}

	// TimeUnit ...
	TimeUnit struct {
		Success int64 `bson:"success" json:"success"`
		Failed  int64 `bson:"failed" json:"failed"`
	}
)

type (
	// SearchKeywordResponse ....
	SearchKeywordResponse struct {
		Keyword    string `json:"keyword"`
		Vietnamese string `json:"vietnamese"`
	}
)

func (k *SearchKeywordRaw) GetPayloadES() *KeywordES {
	return &KeywordES{
		ID:         k.ID.Hex(),
		Vietnamese: k.Vietnamese,
		Keyword:    k.Keyword,
		TimeUnit:   k.TimeUnit,
		Score:      k.Score,
		Source:     k.Source,
		CreatedAt:  k.CreatedAt.Format(config.DateISOFormat),
	}
}
