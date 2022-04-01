package model

type OrderES struct {
	ID                  string                 `json:"id"`
	Code                string                 `json:"code"`
	Status              string                 `json:"status"`
	SearchString        string                 `json:"searchString"`
	ServiceDelivery     string                 `json:"serviceDelivery"`
	Address             string                 `json:"address"`
	IsDeleted           bool                   `json:"isDeleted"`
	FromNewActiveSeller bool                   `json:"fromNewActiveSeller"`
	ApprovedAt          string                 `json:"approvedAt"`
	RejectedAt          string                 `json:"rejectedAt"`
	Price               StatisticPriceES       `json:"price"`
	DeliveredAt         string                 `json:"deliveredAt"`
	CashbackAt          string                 `json:"cashbackAt"`
	CreatedAt           string                 `json:"createdAt"`
	User                string                 `json:"user"`
	Customer            string                 `json:"customer"`
	Remarks             string                 `json:"remarks,omitempty"`
	Delivery            OrderDeliveryES        `json:"delivery"`
	InventoryID         string                 `json:"inventoryId"`
	InventoryName       string                 `json:"inventoryName"`
	SupplierID          string                 `json:"supplierId"`
	SupplierName        string                 `json:"supplierName"`
	Payment             OrderPaymentES         `json:"payment"`
	Location            string                 `json:"location"`
	Banned              bool                   `json:"banned"`
	Wholesale           bool                   `json:"wholesale"`
	TeamMemberID        string                 `json:"teamMemberId,omitempty"`
	TeamID              string                 `json:"teamId,omitempty"`
	InvoiceID           string                 `json:"invoiceId,omitempty"`
	FromNewActiveBuyer  bool                   `json:"fromNewActiveBuyer"`
	Tags                []TagInfo              `json:"tags"`
	Source              string                 `json:"source"`
	IsWaitingCancelled  bool                   `json:"isWaitingCancelled"`
	SendEmailStatus     string                 `json:"sendEmailStatus"`
	MerchantStatus      string                 `json:"merchantStatus"`
	IsAutoApproved      bool                   `json:"isAutoApproved"`
	IsCalled            bool                   `json:"isCalled"`
	ProcessStatus       string                 `json:"processStatus"`
	OutboundRequest     OrderOutboundRequestES `json:"outboundRequest"`
	IsReviewed          bool                   `json:"isReviewed"`
	IsPreorder          bool                   `json:"isPreorder"`
}

type TagInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrderDeliveryES struct {
	SourceID        int    `json:"sourceId"`
	Code            string `json:"code"`
	Status          string `json:"status"`
	ServiceDelivery string `json:"serviceDelivery"`
}

type CartInventoryES struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SKUCommonInfoES ...
type SKUCommonInfoES struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrderPaymentES ...
type OrderPaymentES struct {
	Method string `json:"method"` // Default: COD
	Status string `json:"status"`
}

// CustomerLocationES ...
type CustomerLocationES struct {
	ID          string `json:"id"`
	FullAddress string `json:"fullAddress,omitempty"`
}

type OrderSendEmailES struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type OrderMerchantES struct {
	Status string `bson:"status" json:"status"`
	Note   string `bson:"note" json:"note"`
}

type OrderOutboundRequestES struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

type StatisticPriceES struct {
	Base  float64 `bson:"base" json:"base"`
	Sell  float64 `bson:"sell" json:"sell"`
	Total float64 `bson:"total" json:"total"`
}
