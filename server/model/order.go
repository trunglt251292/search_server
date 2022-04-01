package model

import (
	"fmt"
	"search_server/server/config"
	"search_server/server/internal/util"
	"strings"
	"time"
)

type (
	// OrderRaw ...
	OrderRaw struct {
		ID                 AppID           `bson:"_id" json:"_id"`
		User               AppID           `bson:"user" json:"user"`
		Customer           AppID           `bson:"customer" json:"customer"`
		UserAgent          *UserAgent      `bson:"userAgent,omitempty" json:"userAgent,omitempty"`
		Remarks            string          `bson:"remarks,omitempty"`
		IsChangeDelivery   bool            `bson:"isChangeDelivery"`
		Delivery           *OrderDelivery  `bson:"delivery" json:"delivery"`
		Date               time.Time       `bson:"date" json:"date"`
		Hour               int             `bson:"hour" json:"hour"`
		SearchString       string          `bson:"searchString" json:"searchString"`
		Inventory          *CartInventory  `bson:"inventory" json:"inventory"`
		Supplier           *SKUCommonInfo  `bson:"supplier,omitempty" json:"supplier,omitempty"`
		Payment            *OrderPayment   `bson:"payment" json:"payment"`
		Code               string          `bson:"code" json:"code"`
		CodeOsiris         string          `bson:"codeOsiris,omitempty" json:"codeOsiris,omitempty"`
		DeliveryCode       string          `bson:"deliveryCode,omitempty" json:"deliveryCode,omitempty"`
		TrackingCode       string          `bson:"trackingCode,omitempty" json:"trackingCode,omitempty"`
		TrackingOrderCode  string          `bson:"trackingOrderCode,omitempty" json:"trackingOrderCode,omitempty"`
		TrackingCodeURL    string          `bson:"trackingCodeURL,omitempty" json:"trackingCodeUrl,omitempty"`
		Status             string          `bson:"status" json:"status"`
		Price              *StatisticPrice `bson:"price" json:"price"`
		EstimateCashbackAt time.Time       `bson:"estimateCashbackAt,omitempty" json:"estimateCashbackAt,omitempty"`
		Banned             bool            `bson:"banned" json:"banned"`
		RequestID          AppID           `bson:"requestId,omitempty" json:"requestId"`
		ApprovedAt         time.Time       `bson:"approvedAt,omitempty" json:"approvedAt"`
		RejectedAt         time.Time       `bson:"rejectedAt,omitempty" json:"rejectedAt"`
		PickupAt           time.Time       `bson:"pickupAt,omitempty" json:"pickupAt"`
		DeliveredAt        time.Time       `bson:"deliveredAt,omitempty" json:"deliveredAt"`
		DeliveringAt       time.Time       `bson:"deliveringAt,omitempty" json:"deliveringAt"`
		CashbackAt         time.Time       `bson:"cashbackAt,omitempty" json:"cashbackAt"`
		CreatedAt          time.Time       `bson:"createdAt" json:"createdAt"`
		UpdatedAt          time.Time       `bson:"updatedAt" json:"updatedAt"`

		Wholesale bool `bson:"wholesale" json:"wholesale"`

		TeamMemberID AppID `bson:"teamMemberId,omitempty" json:"teamMemberId,omitempty"`

		TeamID AppID `bson:"teamId,omitempty" json:"teamId,omitempty"`

		InvoiceID AppID `bson:"invoiceId,omitempty" json:"invoiceId,omitempty"`

		Note string `bson:"note" json:"note"`

		Items interface{} `bson:"-" json:"items"`

		StaffApprove AppID `bson:"staffApprove,omitempty" json:"-"`

		Reason          string `bson:"reason" json:"reason"`
		MembershipOrder AppID  `bson:"membership,omitempty"`

		FromNewActiveSeller bool `bson:"fromNewActiveSeller"`

		FromNewActiveBuyer bool `bson:"fromNewActiveBuyer"`

		IsDeleted bool `bson:"isDeleted" json:"isDeleted"`

		Tags []AppID `bson:"tags"`

		RefundTransaction AppID `bson:"refundTransaction,omitempty" json:"refundTransaction,omitempty"`
		TeamOrderBonus    AppID `bson:"teamBonus,omitempty"`

		TotalItem int64  `bson:"totalItem"`
		Source    string `bson:"source" json:"source"`
		Shop      AppID  `bson:"shop,omitempty" json:"shop"`

		IsAssignCoupon bool `bson:"isAssignCoupon" json:"isAssignCoupon"`

		IsWaitingCancelled  bool   `bson:"isWaitingCancelled" json:"isWaitingCancelled"`
		WaitingCancelReason string `bson:"waitingCancelReason,omitempty" json:"waitingCancelReason"`
		WaitingCancelBy     string `bson:"waitingCancelBy,omitempty" json:"waitingCancelBy"`

		HookTimeLastAt time.Time `bson:"hookTimeLastAt,omitempty" json:"hookTimeLastAt,omitempty"`

		SendEmail OrderSendEmail `bson:"sendEmail" json:"sendEmail"`
		Merchant  *OrderMerchant `bson:"merchant,omitempty" json:"merchant,omitempty"`

		IsAutoApproved bool `bson:"isAutoApproved" json:"isAutoApproved"`
		IsCalled       bool `bson:"isCalled" json:"isCalled"`

		ProcessStatus   string                `bson:"processStatus" json:"processStatus"`
		OutboundRequest *OrderOutboundRequest `bson:"outboundRequest" json:"outboundRequest"`
		IsReviewed      bool                  `bson:"isReviewed" json:"isReviewed"`
		IsPreOrder      bool                  `bson:"isPreorder" json:"isPreorder"`
		RestockAt       time.Time             `bson:"restockAt,omitempty" json:"restockAt,omitempty"`
	}

	CartInventory struct {
		ID                    AppID   `json:"_id" bson:"_id"`
		Code                  int     `bson:"id,omitempty" json:"id,omitempty"`
		CanIssueInvoice       bool    `bson:"canIssueInvoice" json:"canIssueInvoice"`
		Name                  string  `bson:"name" json:"name"`
		MinimumValue          float64 `bson:"-" json:"minimumValue"`
		InvoiceDeliveryMethod string  `bson:"-" json:"invoiceDeliveryMethod"`
	}

	// OrderWeight ...
	OrderWeight struct {
		Real                    float64 `json:"real" bson:"real"`
		Converted               float64 `json:"converted" bson:"converted"`
		ConvertedWithMultiplier float64 `json:"convertedWithMultiplier" bson:"convertedWithMultiplier"`
		Multiplier              float64 `json:"multiplier" bson:"multiplier"`
		Sent                    float64 `json:"sent" bson:"sent"`
	}

	// OrderMerchant ...
	OrderMerchant struct {
		Status    string    `bson:"status" json:"status"`
		Note      string    `bson:"note" json:"note"`
		Tags      []AppID   `bson:"tags" json:"tags"`
		UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	}

	// OrderOutboundRequest ...
	OrderOutboundRequest struct {
		ID                  AppID  `bson:"_id" json:"_id"`
		Code                string `bson:"code" json:"code"`
		RequestID           int    `bson:"requestId" json:"requestId"`
		Status              string `bson:"status" json:"status"`
		Reason              string `bson:"reason,omitempty" json:"reason"`
		ErrorCode           string `bson:"errorCode,omitempty" json:"errorCode"`
		PartnerIdentityCode string `bson:"partnerIdentityCode" json:"partnerIdentityCode"`
	}

	// UserAgent ...
	UserAgent struct {
		IP       string `bson:"ip" json:"ip"`
		Version  string `bson:"version" json:"version"`
		Platform string `bson:"platform" json:"platform"`
	}

	// OrderAnalytic ...
	OrderAnalytic struct {
		ID          AppID    `bson:"_id" json:"_id"`
		User        AppID    `bson:"seller"`
		Customer    AppID    `bson:"buyer"`
		Address     string   `bson:"address"`
		Status      string   `bson:"status" json:"status"`
		TotalAmount float64  `bson:"totalAmount"`
		SKUItems    []string `bson:"skuItems"`
		Inventory   string   `bson:"inventory"`

		ApprovedAt  time.Time `bson:"approvedAt,omitempty"`
		RejectedAt  time.Time `bson:"rejectedAt,omitempty"`
		PickupAt    time.Time `bson:"pickupAt,omitempty"`
		DeliveredAt time.Time `bson:"deliveredAt,omitempty"`

		CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
		UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	}

	// StatisticPrice ...
	StatisticPrice struct {
		Base                   float64 `bson:"base" json:"base"`
		Sell                   float64 `bson:"sell" json:"sell"`
		Profit                 float64 `bson:"profit" json:"profit"`
		ProfitUni              float64 `bson:"profitUni" json:"profitUni"`
		Minimum                float64 `bson:"minimum" json:"minimum"`
		Maximum                float64 `bson:"maximum" json:"maximum"`
		Weight                 float64 `bson:"weight" json:"weight"`
		SupplierPrice          float64 `bson:"supplierPrice" json:"supplierPrice"`
		Market                 float64 `bson:"market,omitempty" json:"market,omitempty"`
		Total                  float64 `bson:"total,omitempty" json:"total,omitempty"`
		Volume                 string  `bson:"volume,omitempty" json:"volume,omitempty"`
		MembershipPercent      float64 `bson:"membershipPercent,omitempty" json:"membershipPercent,omitempty"`
		MembershipCommission   float64 `bson:"membershipCommission,omitempty" json:"membershipCommission,omitempty"`
		Promotion              float64 `bson:"promotion,omitempty" json:"promotion,omitempty"`
		TeamBonusPercent       float64 `bson:"teamBonusPercent" json:"teamBonusPercent"`
		TeamBonusCommission    float64 `bson:"teamBonusCommission" json:"teamBonusCommission"`
		WholesaleBonus         float64 `bson:"wholesaleBonus" json:"wholesaleBonus"`
		SellyWholesaleBonus    float64 `bson:"sellyWholesaleBonus" json:"sellyWholesaleBonus"`
		VoucherBonusCommission float64 `bson:"voucherBonusCommission" json:"voucherBonusCommission"`
	}

	// OrderPayment ...
	OrderPayment struct {
		ID                  AppID  `bson:"_id,omitempty" json:"_id,omitempty"`
		Method              string `bson:"method" json:"method"` // Default: COD
		Status              string `bson:"status" json:"status"`
		IsCompletedBySystem bool   `bson:"isCompletedBySystem" json:"isCompletedBySystem"`
	}

	// OrderDelivery ...
	OrderDelivery struct {
		ID                   AppID               `json:"_id" bson:"_id"`
		Status               string              `json:"status" bson:"status"`
		SourceID             int                 `json:"sourceId" bson:"sourceId"`
		Title                string              `bson:"title" json:"title"`
		Code                 string              `bson:"code,omitempty" json:"code,omitempty"`
		FreeShip             bool                `bson:"freeShip" json:"-"`
		Fee                  float64             `bson:"fee" json:"fee"`
		RealShippingFee      float64             `bson:"realShippingFee" json:"realShippingFee"`
		FeeShip              *FeeShippingService `bson:"feeShip,omitempty" json:"feeShip,omitempty"`
		DiscountValue        float64             `bson:"discountValue" json:"discountValue"`
		Note                 string              `bson:"note" json:"note"`
		CourierName          string              `bson:"courierName" json:"courierName"`
		ServiceName          string              `bson:"serviceName" json:"serviceName"`
		ServiceCode          string              `bson:"serviceCode" json:"serviceCode"`
		Currency             string              `bson:"currency" json:"currency"`
		ServiceDelivery      string              `bson:"serviceDelivery" json:"serviceDelivery"` // Unibag or Osiris
		Location             CustomerLocation    `bson:"location" json:"location"`
		CustomerName         string              `bson:"customerName" json:"customerName"`
		CustomerPhone        string              `bson:"customerPhone" json:"customerPhone"`
		CustomerEmail        string              `bson:"customerEmail" json:"customerEmail"`
		MinDeliveryDay       float64             `bson:"minDeliveryDay" json:"minDeliveryDay"`
		MaxDeliveryDay       float64             `bson:"maxDeliveryDay" json:"maxDeliveryDay"`
		EstimateDeliveryAt   time.Time           `bson:"estimateDeliveryAt,omitempty" json:"estimateDeliveryAt,omitempty"` // Now() + min
		EstimateTimeDelivery string              `bson:"estimateTimeDelivery,omitempty" json:"estimate_time_delivery"`
		Reason               string              `bson:"reason" json:"reason"`

		DeliveryCode string       `bson:"-" json:"deliveryCode,omitempty"`
		TrackingURL  string       `json:"trackingUrl,omitempty" bson:"-"`
		OrderWeight  *OrderWeight `bson:"orderWeight,omitempty" json:"orderWeight"`
	}

	CustomerLocation struct {
		ID          AppID  `bson:"_id" json:"_id,omitempty"`
		Province    int    `bson:"province" json:"province"`
		District    int    `bson:"district" json:"district"`
		Ward        int    `bson:"ward" json:"ward"`
		Address     string `bson:"address" json:"address"`
		FullAddress string `bson:"fullAddress,omitempty" json:"fullAddress,omitempty"`
	}

	// OrderResponseORS ...
	OrderResponseORS struct {
		Success   bool    `json:"success"`
		OrderCode string  `json:"order_code"`
		Revenue   float64 `json:"revenue"`
	}

	// OrderResponseIMS ...
	OrderResponseIMS struct {
		Data    OrderIMS `json:"data"`
		Code    int      `json:"code"`
		Message string   `json:"message"`
	}

	// OrderResponseIMSOrder ...
	OrderResponseIMSOrder struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Success bool   `json:"success"`
	}

	// OrderIMS ...
	OrderIMS struct {
		ShippingFee  float64      `bson:"shippingFee" json:"shippingFee"`
		TrackingCode string       `bson:"trackingCode,omitempty" json:"trackingCode,omitempty"`
		TrackingLink string       `bson:"trackingLink,omitempty" json:"trackingLink,omitempty"`
		Code         string       `bson:"orderCode,omitempty" json:"orderCode,omitempty"`
		Status       string       `bson:"status" json:"status"`
		ServiceName  string       `bson:"serviceName" json:"serviceName"`
		ServiceCode  string       `bson:"serviceCode" json:"serviceCode"`
		CourierName  string       `bson:"courierName" json:"courierName"`
		SourceID     int          `bson:"sourceId" json:"sourceId"`
		SourceName   string       `bson:"sourceName,omitempty" json:"sourceName,omitempty"`
		Weight       *OrderWeight `bson:"weight" json:"weight"`
	}

	OrderSendEmail struct {
		ID     *AppID `bson:"_id,omitempty" json:"_id,omitempty"`
		Status string `json:"status" bson:"status"`
	}

	// OrderTeamAnalytics ...
	OrderTeamAnalytics struct {
		User         AppID     `json:"user" bson:"user"`
		DeliveredAt  time.Time `json:"deliveredAt" bson:"deliveredAt"`
		TeamMemberID AppID     `bson:"teamMemberId,omitempty" json:"teamMemberId,omitempty"`
		TeamID       AppID     `bson:"teamId,omitempty" json:"teamId,omitempty"`
		Status       string    `bson:"status" json:"status"`
	}

	// OrderAggregateIDByInventory ...
	OrderAggregateIDByInventory struct {
		InventoryID AppID   `bson:"_id"`
		OrderIDs    []AppID `bson:"orderIds"`
	}

	FeeShippingService struct {
		Total     float64 `bson:"total" json:"total"`
		Insurance float64 `bson:"insurance" json:"insurance"`
		COD       float64 `bson:"cod" json:"cod"`
		Shipping  float64 `bson:"shipping" json:"shipping"`
	}
)

func (o *OrderRaw) GetSearchStringES(userPhone string, skus []string) string {
	phone := util.PhoneNumberFormatFromPhone(o.Delivery.CustomerPhone)
	userPhoneFormat84 := fmt.Sprintf("84%s", userPhone)
	userPhoneFormat0 := fmt.Sprintf("0%s", userPhone)

	return fmt.Sprintf("%s %s %s %s %s %s %s %s %s %s %s", o.Code, o.DeliveryCode, o.Delivery.Location.FullAddress, o.Delivery.CustomerName, o.Delivery.CustomerPhone, "0"+phone.Number, userPhoneFormat84, userPhoneFormat0, strings.Join(skus, " "), o.Remarks, o.Reason)
}

func (o *OrderRaw) GetPayloadES(userPhone string, skus []string, tags []TagInfo) *OrderES {
	res := &OrderES{
		ID:                  o.ID.Hex(),
		Code:                o.Code,
		Status:              o.Status,
		SearchString:        o.GetSearchStringES(userPhone, skus),
		ServiceDelivery:     o.Delivery.ServiceDelivery,
		Address:             o.Delivery.Location.FullAddress,
		IsDeleted:           o.IsDeleted,
		FromNewActiveSeller: o.FromNewActiveSeller,
		ApprovedAt:          o.ApprovedAt.Format(config.DateISOFormat),
		RejectedAt:          o.RejectedAt.Format(config.DateISOFormat),
		Price: StatisticPriceES{
			Base:  o.Price.Base,
			Sell:  o.Price.Sell,
			Total: o.Price.Total,
		},
		DeliveredAt: o.DeliveredAt.Format(config.DateISOFormat),
		CashbackAt:  o.CashbackAt.Format(config.DateISOFormat),
		CreatedAt:   o.CreatedAt.Format(config.DateISOFormat),
		User:        o.User.Hex(),
		Customer:    o.Customer.Hex(),
		Remarks:     o.Remarks,
		Delivery: OrderDeliveryES{
			SourceID:        o.Delivery.SourceID,
			Code:            o.Delivery.Code,
			Status:          o.Delivery.Status,
			ServiceDelivery: o.Delivery.ServiceDelivery,
		},
		Payment:            OrderPaymentES{},
		Banned:             o.Banned,
		Wholesale:          o.Wholesale,
		FromNewActiveBuyer: o.FromNewActiveBuyer,
		Tags:               tags,
		Source:             o.Source,
		IsWaitingCancelled: o.IsWaitingCancelled,
		SendEmailStatus:    o.SendEmail.Status,
		IsAutoApproved:     o.IsAutoApproved,
		IsCalled:           o.IsCalled,
		ProcessStatus:      o.ProcessStatus,
		IsReviewed:         o.IsReviewed,
		IsPreorder:         o.IsPreOrder,
	}
	if o.Payment != nil {
		res.Payment = OrderPaymentES{
			Method: o.Payment.Method,
			Status: o.Payment.Status,
		}
	}
	if o.Inventory != nil {
		res.InventoryID = o.Inventory.ID.Hex()
		res.InventoryName = o.Inventory.Name
	}
	if o.Supplier != nil {
		res.SupplierID = o.Supplier.ID.Hex()
		res.SupplierName = o.Supplier.Name
	}
	if !o.TeamMemberID.IsZero() {
		res.TeamMemberID = o.TeamMemberID.Hex()
		res.TeamID = o.TeamID.Hex()
	}
	if !o.InvoiceID.IsZero() {
		res.InvoiceID = o.InvoiceID.Hex()
	}
	if o.Merchant != nil {
		res.MerchantStatus = o.Merchant.Status
	}
	if o.OutboundRequest != nil {
		res.OutboundRequest = OrderOutboundRequestES{
			Code:   o.OutboundRequest.Code,
			Status: o.OutboundRequest.Status,
		}
	}
	return res
}
