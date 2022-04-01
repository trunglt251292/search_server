package model

import (
	"search_server/server/config"
	"strings"
	"time"
)

// UserRaw ...
type UserRaw struct {
	ID                     AppID                 `bson:"_id" json:"_id"`
	Name                   string                `bson:"name" json:"name"`
	SearchString           string                `bson:"searchString" json:"-"`
	CreatedAt              time.Time             `bson:"createdAt" json:"createdAt"`
	UpdatedAt              time.Time             `bson:"updatedAt" json:"updatedAt"`
	Phone                  *Phone                `bson:"phone" json:"phone"`
	Info                   UserContactInfo       `bson:"info" json:"info"`
	Banned                 bool                  `bson:"banned" json:"banned"`
	HasOrder               bool                  `bson:"hasOrder" json:"hasOrder"`
	BannedReason           string                `bson:"bannedReason,omitempty"`
	LastActivatedAt        time.Time             `bson:"lastActivatedAt" json:"lastActivatedAt"`
	LastViewNotificationAt time.Time             `bson:"lastViewNotificationAt" json:"lastViewNotificationAt"`
	RegisterFrom           string                `bson:"registerFrom" json:"registerFrom"`
	Code                   string                `bson:"code" json:"code"`
	IsUpdatedInfo          bool                  `bson:"isUpdatedInfo" json:"isUpdatedInfo"`
	Referral               *ReferralInfo         `bson:"referral" json:"referral"`
	Membership             MembershipInfo        `bson:"membership" json:"membership"`
	Identification         *IdentificationUser   `bson:"identification,omitempty" json:"identification,omitempty"`
	Team                   *UserTeamInfo         `bson:"team,omitempty" json:"team,omitempty"`
	Segment                string                `bson:"segment"`
	Apple                  *UserAppleData        `json:"apple,omitempty" bson:"apple,omitempty"`
	Facebook               *UserFacebookData     `json:"facebook,omitempty" bson:"facebook,omitempty"`
	Google                 *UserGoogleData       `json:"google,omitempty" bson:"google,omitempty"`
	SocialLoginEmail       *UserSocialLoginEmail `json:"email,omitempty" bson:"email,omitempty"`
	Segments               []AppID               `bson:"segments,omitempty"`
	Zalo                   *UserZaloData         `json:"zalo,omitempty" bson:"zalo,omitempty"`
}

// ReferralInfo ...
type ReferralInfo struct {
	Code         string    `json:"code" bson:"code"`
	Enabled      bool      `json:"enabled" bson:"enabled"`
	EnabledAt    time.Time `json:"enabledAt" bson:"enabledAt"`
	ShareContent string    `json:"shareContent" bson:"shareContent"`
	TotalInvitee int64     `json:"totalInvitee" bson:"-"`
}

// UserContactInfo ...
type UserContactInfo struct {
	Email    string `bson:"email" json:"email"`
	City     int    `bson:"cityCode" json:"cityCode"`
	CityName string `bson:"-" json:"cityName"`
	Gender   string `bson:"gender" json:"gender"`
}

// Phone ...
type Phone struct {
	CountryCode string    `json:"countryCode" bson:"countryCode"`
	Number      string    `json:"number" bson:"number"`
	Full        string    `json:"full" bson:"full"`
	Verified    bool      `json:"verified" bson:"verified"`
	VerifiedAt  time.Time `json:"verifiedAt" bson:"verifiedAt,omitempty"`
}

// UserZaloData ...
type UserZaloData struct {
	ID    string `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Token string `bson:"token" json:"-"`
	Photo string `bson:"photo" json:"photo"`
}

// UserSocialLoginEmail ...
type UserSocialLoginEmail struct {
	Email      string    `bson:"email" json:"email"`
	Verified   bool      `bson:"verified" json:"verified"`
	VerifiedAt time.Time `bson:"verifiedAt" json:"verifiedAt"`
}

// UserAppleData ....
type UserAppleData struct {
	ID    string `bson:"id"`
	Email string `bson:"email"`
	Name  string `bson:"name"`
}

// UserFacebookData ...
type UserFacebookData struct {
	ID     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
	Token  string `json:"token" bson:"token"`
	Photo  string `json:"photo" bson:"photo"`
	Gender string `json:"gender" bson:"gender"`
}

// UserGoogleData ...
type UserGoogleData struct {
	ID    string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
	Name  string `json:"name" bson:"name"`
	Photo string `json:"photo" bson:"photo"`
	Token string `json:"token" bson:"token"`
}

// IdentificationUser ...
type IdentificationUser struct {
	ID          AppID     `bson:"_id" json:"_id"`
	Status      string    `bson:"status" json:"status"`
	TimeUpdate  int       `bson:"timeUpdate" json:"timeUpdate"`
	CompletedAt time.Time `bson:"completedAt,omitempty" json:"completedAt,omitempty"`
	Note        string    `bson:"note,omitempty" json:"note,omitempty"`
}

// MembershipInfo ...
type MembershipInfo struct {
	CurrentLevel            int       `bson:"currentLevel" `
	CurrentTransactionCount int       `bson:"currentTransactionCount" `
	ExpireAt                time.Time `bson:"expireAt"`
	NextLevel               int       `bson:"nextLevel"`
	CurrentSales            float64   `bson:"currentSales"`
}

// UserTeamInfo ...
type UserTeamInfo struct {
	ID    AppID          `bson:"_id" json:"_id"`
	Name  string         `bson:"name" json:"name"`
	Role  string         `bson:"role" json:"role"`
	Level *TeamLevelInfo `bson:"-" json:"level,omitempty"`
}

// TeamLevelInfo ...
type TeamLevelInfo struct {
	Name                    string    `bson:"name" json:"name"`
	Color                   string    `bson:"color" json:"color"`
	Level                   int       `bson:"level" json:"level"`
	LevelAtStartOfMonth     int       `bson:"levelAtStartOfMonth" json:"levelAtStartOfMonth"`
	NextLevel               int       `bson:"nextLevel" json:"nextLevel"`
	ExpiredAt               time.Time `bson:"expiredAt" json:"expiredAt"`
	MaximumNumberMember     int       `bson:"maximumNumberMember" json:"maximumNumberMember"`
	MaximumVipMember        int       `bson:"maximumVipMember" json:"maximumVipMember"`
	MaximumNumberViceLeader int       `bson:"maximumNumberViceLeader" json:"maximumNumberViceLeader"`
	BonusPercent            float64   `bson:"bonusPercent" json:"bonusPercent"`
}

func (p *Phone) GetSearchString() string {
	return p.Full + " " + p.Number + " " + strings.Replace(p.Full, "+84", "0", 1)
}

// GetSearchStringES ...
func (u *UserRaw) GetSearchStringES() string {
	searchText := u.Name + " " + u.Info.Email + " " + u.Code
	if u.Phone != nil {
		searchText += " " + u.Phone.GetSearchString()
	}
	if u.Facebook != nil {
		searchText += " " + u.Facebook.ID
	}
	if u.Google != nil {
		searchText += " " + u.Google.ID
	}
	if u.Apple != nil {
		searchText += " " + u.Apple.ID
	}
	if u.SocialLoginEmail != nil {
		searchText += " " + u.SocialLoginEmail.Email
	}
	return searchText
}

// GetPayloadES ...
func (u *UserRaw) GetPayloadES(inviter AppID) *UserES {
	res := &UserES{
		ID:   u.ID.Hex(),
		Name: u.Name,
		Info: Info{
			CityCode: u.Info.City,
			Gender:   u.Info.Gender,
			Email:    u.Info.Email,
		},
		MembershipLevel:     u.Membership.CurrentLevel,
		MembershipNextLevel: u.Membership.NextLevel,
		SearchString:        u.GetSearchStringES(),
		Code:                u.Code,
		HasOrder:            u.HasOrder,
		LastActivatedAt:     u.LastActivatedAt.Format(config.DateISOFormat),
		Banned:              u.Banned,
		Invitee:             "",
		Segments:            make([]SegmentES, 0),
		CreatedAt:           u.CreatedAt.Format(config.DateISOFormat),
	}
	if !inviter.IsZero() {
		res.Invitee = inviter.Hex()
	}
	if len(u.Segments) > 0 {
		for _, sg := range u.Segments {
			res.Segments = append(res.Segments, SegmentES{ID: sg.Hex()})
		}
	}
	if u.Phone != nil {
		res.Phone = PhoneUserES{
			CountryCode: u.Phone.CountryCode,
			Number:      u.Phone.Number,
			Full:        u.Phone.Full,
			Verified:    u.Phone.Verified,
		}
	}
	return res
}
