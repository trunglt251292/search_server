package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/Selly-Modules/elasticsearch"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
)

// CommonQuery ...
type CommonQuery struct {
	ClearCashRequestID string

	IsOutOfStock    string
	CanIssueInvoice string

	EventTicket AppID
	IsPreorder  string

	Score       int
	Transaction string
	Staff       string
	Device      string
	IDFA        string // Adjust idad
	ADID        string // Adjust adid

	HasContent string
	HasMedia   string

	ProcessStatus         string
	OutboundRequestStatus string

	IsAutoApproved string

	IsCalled string

	DeliveryStatus string

	TimeRange string

	EventRewardType string

	TransferID      string
	OsName          string
	AppVersion      string
	OsVersion       string
	LastActivatedAt time.Time

	LeaderboardID string

	Brands    []string
	Cities    []string
	Suppliers []string

	SourceDelivery string

	Tags      string
	IsPin     string
	Order     int
	PageToken string
	CreatedBy string // "me" or id hex ...
	InvitedBy string // inviter id

	Region              string
	IsDeleted           string
	CityCode            string
	CitySlug            string
	ForceSendSms        string
	Action              string
	Apply               string
	PaymentMethod       string
	OS                  string
	FromNewActiveSeller string
	VersionCode         int64
	Page                int64
	Limit               int64
	Sort                bson.D
	IsLeft              string
	SortStr             string
	User                string
	NotUsers            string
	RoundID             string
	Categories          string
	SubCategory         string
	InventoryID         string
	NotInventory        string
	ListNotInventory    string
	NotInventories      []string
	Product             string
	Keyword             string
	Timestamp           time.Time
	Status              string
	Success             string
	Service             string
	Active              string
	Type                string
	Category            string
	CityID              string
	DistrictID          string
	ServiceDelivery     string
	Banned              string
	FieldSearch         string
	FromAt              time.Time
	ToAt                time.Time

	Slug string

	DeliveredFrom time.Time
	DeliveredTo   time.Time
	CashbackFrom  time.Time
	CashbackTo    time.Time
	ApprovedFrom  time.Time
	ApprovedTo    time.Time

	FromAtUnix                int64
	ToAtUnix                  int64
	FromPrice                 float64
	ToPrice                   float64
	IgnoreIDs                 string // "id1,id2,..."
	Supplier                  string
	Inventory                 string
	ListInventory             string
	Author                    string
	CardNumber                string
	IncludeDeclined           string
	Declined                  []AppID
	Group                     string
	IDs                       []AppID
	IsQueryInIDs              bool
	Types                     string
	Time                      string
	Source                    string
	UpdatedAt                 int64
	CategoryId                string
	Collection                string
	Brand                     string
	HagTag                    string
	Timer                     string
	Published                 string
	FromSystem                string
	MembershipLevel           string
	NoBrand                   string
	IsNotStatusCancelled      bool
	Branch                    string
	Bank                      string
	TargetType                string
	TargetValue               string
	Code                      string
	Sku                       string
	Products                  string
	HasUpdate                 string
	ProvinceCode              int
	PendingInactive           string
	Name                      string
	Emails                    []string
	IncludeQuest              string
	Recipient                 string
	Seller                    string
	InvoiceRequestID          string
	IncludePriceSell          bool
	StatusApprovedFromAt      time.Time
	StatusApprovedToAt        time.Time
	City                      string
	HasAmount                 string
	AdminNotificationCategory string
	EmailStatus               string
	MerchantStatus            string
	IsPresident               string
	Level                     string
	IsFull                    string
	SKUIDs                    []AppID
	EventQuest                string
	Statuses                  string
	Users                     string
	IsWholesaleBonus          string
	Segments                  string
	DistrictSlug              string
	PromotionID               string
}

// ConvertESQuery ...
func (q *CommonQuery) ConvertESQuery(indexName string) *elasticsearch.ESQuery {
	qES := &elasticsearch.ESQuery{
		Index: indexName,
		Sorts: []elasticsearch.ESSort{
			{
				Field:     "_score",
				Ascending: false,
			},
		},
	}

	if len(q.Sort) > 0 {
		for _, sort := range q.Sort {
			sortES := elasticsearch.ESSort{
				Field:     sort.Key,
				Ascending: true,
			}
			if sort.Value.(int) == -1 {
				sortES.Ascending = false
			}
			qES.Sorts = append(qES.Sorts, sortES)
		}
	}

	// Page
	if q.Page > 0 {
		qES.Page = q.Page
	}

	// Limit
	if q.Limit > 0 {
		qES.Limit = q.Limit
	}
	// Keyword
	if len(q.Keyword) > 0 {
		qES.Keyword = q.Keyword
	}

	if !q.FromAt.IsZero() {
		qES.FromAt = q.FromAt
	}

	if !q.ToAt.IsZero() {
		qES.ToAt = q.ToAt
	}

	if !q.CashbackFrom.IsZero() {
		qES.CashbackFrom = q.CashbackFrom
	}

	if !q.CashbackTo.IsZero() {
		qES.CashbackTo = q.CashbackTo
	}

	if !q.DeliveredFrom.IsZero() {
		qES.DeliveredFrom = q.DeliveredFrom
	}

	if !q.DeliveredTo.IsZero() {
		qES.DeliveredTo = q.DeliveredTo
	}

	if !q.ApprovedFrom.IsZero() {
		qES.ApprovedFrom = q.ApprovedFrom
	}

	if !q.ApprovedTo.IsZero() {
		qES.ApprovedTo = q.ApprovedTo
	}

	if q.FromPrice > 0 {
		qES.FromPrice = q.FromPrice
	}

	if q.ToPrice > 0 {
		qES.ToPrice = q.ToPrice
	}
	q.AssignBrandsES(qES)
	q.AssignUserES(qES)
	q.AssignBannedES(qES)
	q.AssignSourceDeliveryES(qES)
	q.AssignServiceDeliveryES(qES)
	q.AssignActiveES(qES)
	q.AssignCategoriesES(qES)
	q.AssignListNotUserES(qES)
	q.AssignSubCategoriesES(qES)
	q.AssignIgnoreIdsES(qES)
	q.AssignSuppliersES(qES)
	q.AssignCitiesES(qES)
	q.AssignTypeES(qES)
	q.AssignSourceES(qES)
	q.AssignFromNewActiveSellerES(qES)
	q.AssignEmailStatusES(qES)
	q.AssignMerchantStatusES(qES)
	q.AssignIsCalledES(qES)
	q.AssignIsOutOfStockES(qES)
	q.AssignPendingInactiveES(qES)
	q.AssignCanIssueInvoiceES(qES)
	q.AssignIsAutoApprovedES(qES)
	q.AssignProcessStatusES(qES)
	q.AssignOutboundRequestStatusES(qES)
	q.AssignIsWholesaleES(qES)
	q.AssignIsPreorderES(qES)
	q.AssignIsDeletedES(qES)
	q.AssignTagsES(qES)
	q.AssignListStatusES(qES)
	q.AssignDeliveryListStatusES(qES)
	q.AssignInventoriesES(qES)
	q.AssignNotInventoriesES(qES)
	q.AssignMembershipLevelES(qES)
	q.AssignInvitedByES(qES)
	q.AssignSegmentsES(qES)

	return qES
}

// AssignSegmentsES ...
func (q *CommonQuery) AssignSegmentsES(qES *elasticsearch.ESQuery) {
	if q.Segments != "" {
		list := strings.Split(q.Segments, ",")
		if len(list) > 0 {
			qES.Segments = list
		}
	}
}

// AssignInvitedByES ...
func (q *CommonQuery) AssignInvitedByES(qES *elasticsearch.ESQuery) {
	if q.InvitedBy != "" {
		qES.Invitee = q.InvitedBy
	}
}

// AssignMembershipLevelES ...
func (q *CommonQuery) AssignMembershipLevelES(qES *elasticsearch.ESQuery) {
	if q.MembershipLevel != "" {
		membership, err := strconv.Atoi(q.MembershipLevel)
		if err == nil {
			qES.MembershipLevel = membership
		}
	}
}

// AssignInventoriesES ...
func (q *CommonQuery) AssignInventoriesES(qES *elasticsearch.ESQuery) {
	if q.InventoryID != "" || q.ListInventory != "" {
		var (
			inventories = make([]string, 0)
		)
		if q.InventoryID != "" {
			inventories = append(inventories, q.InventoryID)
		}
		if q.ListInventory != "" {
			list := strings.Split(q.ListInventory, ",")
			if len(list) > 0 {
				inventories = append(inventories, list...)
			}
		}
		qES.Inventories = inventories
	}
}

// AssignBrandsES ...
func (q *CommonQuery) AssignBrandsES(qES *elasticsearch.ESQuery) {
	if q.Brand != "" || len(q.Brands) > 0 {
		var (
			brands = make([]string, 0)
		)
		if q.Brand != "" {
			brands = append(brands, q.Brand)
		}
		if len(q.Brands) > 0 {
			for _, c := range q.Brands {
				if c != "" {
					brands = append(brands, c)
				}
			}
		}
		qES.Brands = brands
	}
}

// AssignNotInventoriesES ...
func (q *CommonQuery) AssignNotInventoriesES(qES *elasticsearch.ESQuery) {
	if q.ListNotInventory != "" {
		list := strings.Split(q.ListNotInventory, ",")
		if len(list) > 0 {
			qES.NotInventories = list
		}
	}
}

// AssignListStatusES ...
func (q *CommonQuery) AssignListStatusES(qES *elasticsearch.ESQuery) {
	if q.Status != "" {
		listStatus := strings.Split(strings.ToLower(q.Status), ",")
		if len(listStatus) > 0 {
			qES.ListStatus = listStatus
		}
	}
}

// AssignDeliveryListStatusES ...
func (q *CommonQuery) AssignDeliveryListStatusES(qES *elasticsearch.ESQuery) {
	if q.DeliveryStatus != "" {
		listStatus := strings.Split(strings.ToLower(q.DeliveryStatus), ",")
		if len(listStatus) > 0 {
			qES.ListDeliveryStatus = listStatus
		}
	}
}

// AssignTagsES ...
func (q *CommonQuery) AssignTagsES(qES *elasticsearch.ESQuery) {
	if q.Tags != "" {
		tags := strings.Split(q.Tags, ",")
		if len(tags) > 0 {
			qES.Tags = tags
		}
	}
}

// AssignIsPreorderES ...
func (q *CommonQuery) AssignIsPreorderES(qES *elasticsearch.ESQuery) {
	if q.IsPreorder != "" && funk.Contains([]string{"true", "false"}, q.IsPreorder) {
		qES.IsPreorder = q.IsPreorder
	}
}

// AssignIsDeletedES ...
func (q *CommonQuery) AssignIsDeletedES(qES *elasticsearch.ESQuery) {
	if q.IsDeleted != "" && funk.Contains([]string{"true", "false"}, q.IsDeleted) {
		qES.IsDeleted = q.IsDeleted
	}
}

// AssignIsCalledES ...
func (q *CommonQuery) AssignIsCalledES(qES *elasticsearch.ESQuery) {
	if q.IsCalled != "" && funk.Contains([]string{"true", "false"}, q.IsCalled) {
		qES.IsCalled = q.IsCalled
	}
}

// AssignIsOutOfStockES ...
func (q *CommonQuery) AssignIsOutOfStockES(qES *elasticsearch.ESQuery) {
	if q.IsOutOfStock != "" && funk.Contains([]string{"true", "false"}, q.IsOutOfStock) {
		qES.IsOutOfStock = q.IsOutOfStock
	}
}

// AssignPendingInactiveES ...
func (q *CommonQuery) AssignPendingInactiveES(qES *elasticsearch.ESQuery) {
	if q.PendingInactive != "" && funk.Contains([]string{"true", "false"}, q.PendingInactive) {
		qES.PendingInactive = q.PendingInactive
	}
}

// AssignCanIssueInvoiceES ...
func (q *CommonQuery) AssignCanIssueInvoiceES(qES *elasticsearch.ESQuery) {
	if q.CanIssueInvoice != "" && funk.Contains([]string{"true", "false"}, q.CanIssueInvoice) {
		qES.CanIssueInvoice = q.CanIssueInvoice
	}
}

// AssignIsWholesaleES ...
func (q *CommonQuery) AssignIsWholesaleES(qES *elasticsearch.ESQuery) {
	if q.IsWholesaleBonus != "" && funk.Contains([]string{"true", "false"}, q.IsWholesaleBonus) {
		qES.IsWholesaleBonus = q.IsWholesaleBonus
	}
}

// AssignIsAutoApprovedES ...
func (q *CommonQuery) AssignIsAutoApprovedES(qES *elasticsearch.ESQuery) {
	if q.IsAutoApproved != "" && funk.Contains([]string{"true", "false"}, q.IsAutoApproved) {
		qES.IsAutoApproved = q.IsAutoApproved
	}
}

// AssignOutboundRequestStatusES ...
func (q *CommonQuery) AssignOutboundRequestStatusES(qES *elasticsearch.ESQuery) {
	if q.OutboundRequestStatus != "" {
		qES.OutboundRequestStatus = strings.ToLower(q.OutboundRequestStatus)
	}
}

// AssignProcessStatusES ...
func (q *CommonQuery) AssignProcessStatusES(qES *elasticsearch.ESQuery) {
	if q.ProcessStatus != "" {
		qES.ProcessStatus = strings.ToLower(q.ProcessStatus)
	}
}

// AssignEmailStatusES ...
func (q *CommonQuery) AssignEmailStatusES(qES *elasticsearch.ESQuery) {
	if q.EmailStatus != "" {
		qES.EmailStatus = strings.ToLower(q.EmailStatus)
	}
}

// AssignMerchantStatusES ...
func (q *CommonQuery) AssignMerchantStatusES(qES *elasticsearch.ESQuery) {
	if q.MerchantStatus != "" {
		qES.MerchantStatus = strings.ToLower(q.MerchantStatus)
	}
}

// AssignFromNewActiveSellerES ...
func (q *CommonQuery) AssignFromNewActiveSellerES(qES *elasticsearch.ESQuery) {
	if q.FromNewActiveSeller != "" && funk.Contains([]string{"true", "false"}, q.FromNewActiveSeller) {
		qES.FromNewActiveSeller = q.FromNewActiveSeller
	}
}

// AssignSourceES ...
func (q *CommonQuery) AssignSourceES(qES *elasticsearch.ESQuery) {
	if q.Source != "" {
		qES.Source = q.Source
	}
}

// AssignPaymentMethodES ...
func (q *CommonQuery) AssignPaymentMethodES(qES *elasticsearch.ESQuery) {
	if q.PaymentMethod != "" {
		qES.PaymentMethod = q.PaymentMethod
	}
}

// AssignTypeES ...
func (q *CommonQuery) AssignTypeES(qES *elasticsearch.ESQuery) {
	if q.Type != "" {
		qES.Type = q.Type
	}
}

// AssignCitiesES ...
func (q *CommonQuery) AssignCitiesES(qES *elasticsearch.ESQuery) {
	if q.CitySlug != "" || len(q.Cities) > 0 {
		var (
			cities = make([]string, 0)
		)
		if q.CitySlug != "" {
			cities = append(cities, q.CitySlug)
		}
		if len(q.Cities) > 0 {
			for _, c := range q.Cities {
				if c != "" {
					cities = append(cities, c)
				}
			}
		}
		qES.SlugCites = cities
	}
}

// AssignSuppliersES ...
func (q *CommonQuery) AssignSuppliersES(qES *elasticsearch.ESQuery) {
	if q.Supplier != "" && len(q.Suppliers) > 0 {
		var (
			suppliers = make([]string, 0)
		)
		if q.Supplier != "" {
			suppliers = append(suppliers, q.Supplier)
		}
		if len(q.Suppliers) > 0 {
			for _, s := range q.Suppliers {
				if s != "" {
					suppliers = append(suppliers, s)
				}
			}
		}
		qES.Suppliers = suppliers
	}
}

// AssignIgnoreIdsES ...
func (q *CommonQuery) AssignIgnoreIdsES(qES *elasticsearch.ESQuery) {
	if q.IgnoreIDs != "" {
		ignoreIds := strings.Split(q.IgnoreIDs, ",")
		if len(ignoreIds) > 0 {
			qES.IgnoreIDs = ignoreIds
		}
	}
}

// AssignSubCategoriesES ...
func (q *CommonQuery) AssignSubCategoriesES(qES *elasticsearch.ESQuery) {
	if q.SubCategory != "" {
		subCategories := strings.Split(q.SubCategory, ",")
		if len(subCategories) > 0 {
			qES.SubCategories = subCategories
		}
	}
}

// AssignListNotUserES ...
func (q *CommonQuery) AssignListNotUserES(qES *elasticsearch.ESQuery) {
	if q.NotUsers != "" {
		notUsers := strings.Split(q.NotUsers, ",")
		if len(notUsers) > 0 {
			qES.ListNotUser = notUsers
		}
	}
}

// AssignUserES ...
func (q *CommonQuery) AssignUserES(qES *elasticsearch.ESQuery) {
	if q.User != "" || q.Users != "" {
		var (
			users = make([]string, 0)
		)
		if q.User != "" {
			users = append(users, q.User)
		}
		if q.Users != "" {
			us := strings.Split(q.Users, ",")
			if len(us) > 0 {
				users = append(users, us...)
			}
		}
		qES.ListUser = users
	}
}

// AssignCategoriesES ...
func (q *CommonQuery) AssignCategoriesES(qES *elasticsearch.ESQuery) {
	if q.Category != "" || q.Categories != "" {
		var (
			categories = make([]string, 0)
		)
		if q.Category != "" {
			categories = append(categories, q.Category)
		}
		if q.Categories != "" {
			ctes := strings.Split(q.Categories, ",")
			if len(ctes) > 0 {
				categories = append(categories, ctes...)
			}
		}
		qES.Categories = categories
	}
}

// AssignBannedES ...
func (q *CommonQuery) AssignBannedES(qES *elasticsearch.ESQuery) {
	if q.Banned != "" && funk.Contains([]string{"true", "false"}, q.Banned) {
		qES.Banned = q.Banned
	}
}

// AssignActiveES ...
func (q *CommonQuery) AssignActiveES(qES *elasticsearch.ESQuery) {
	if q.Active != "" && funk.Contains([]string{"true", "false"}, q.Active) {
		qES.Active = q.Active
	}
}

// AssignSourceDeliveryES ...
func (q *CommonQuery) AssignSourceDeliveryES(qES *elasticsearch.ESQuery) {
	if q.SourceDelivery != "" {
		qES.SourceDelivery = q.SourceDelivery
	}
}

// AssignServiceDeliveryES ...
func (q *CommonQuery) AssignServiceDeliveryES(qES *elasticsearch.ESQuery) {
	if q.ServiceDelivery != "" {
		qES.ServiceDelivery = q.ServiceDelivery
	}
}
