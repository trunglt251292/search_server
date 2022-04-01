package model

// KeywordES ...
type KeywordES struct {
	ID         string   `json:"id"`
	Vietnamese string   `json:"vietnamese"`
	Keyword    string   `json:"keyword"`
	TimeUnit   TimeUnit `json:"timeUnit"`
	Score      int64    `json:"score"`
	Source     string   `json:"source"`
	CreatedAt  string   `json:"createdAt"`
}
