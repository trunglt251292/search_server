package model

type UserES struct {
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	Phone               PhoneUserES `json:"phone"`
	Info                Info        `json:"info"`
	SearchString        string      `json:"searchString"`
	Code                string      `json:"code"`
	HasOrder            bool        `json:"hasOrder"`
	LastActivatedAt     string      `json:"lastActivatedAt"`
	Banned              bool        `json:"banned"`
	MembershipLevel     int         `json:"membershipLevel"`
	MembershipNextLevel int         `json:"membershipNextLevel"`
	Invitee             string      `json:"invitee"`
	Segments            []SegmentES `json:"segments"`
	CreatedAt           string      `json:"createdAt"`
}

type SegmentES struct {
	ID string `json:"id"`
}

type PhoneUserES struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
	Full        string `json:"full"`
	Verified    bool   `json:"verified"`
}

type Info struct {
	CityCode int    `json:"cityCode"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}
