package tnc

// OutboundRequestRes ...
type OutboundRequestRes struct {
	OrID          int    `json:"orId"`
	OrCode        string `json:"orCode"`
	PartnerORCode string `json:"partnerORCode"`
	Error         *Error `json:"error"`
}

type authRes struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

// Webhook ...
type Webhook struct {
	OrId          int    `json:"orId"`
	PartnerORCode string `json:"partnerORCode"`
	ErrorCode     string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
	Event         string `json:"event"`
	Id            string `json:"id"`
	Timestamp     int64  `json:"timestamp"`
	Note          string `json:"note"`
	OrCode        string `json:"orCode"`
}

// OutboundRequestInfo ...
type OutboundRequestInfo struct {
	OrId                  int     `json:"orId"`
	OrCode                string  `json:"orCode"`
	PartnerORCode         string  `json:"partnerORCode"`
	OriginalPartnerOrCode string  `json:"originalPartnerOrCode"`
	PartnerRefId          string  `json:"partnerRefId"`
	RefCode               string  `json:"refCode"`
	WarehouseCode         string  `json:"warehouseCode"`
	Status                string  `json:"status"`
	Note                  string  `json:"note"`
	ShippingType          int     `json:"shippingType"`
	PriorityType          int     `json:"priorityType"`
	PackType              int     `json:"packType"`
	BizType               int     `json:"bizType"`
	CustomerName          string  `json:"customerName"`
	CustomerPhoneNumber   string  `json:"customerPhoneNumber"`
	ShippingFullAddress   string  `json:"shippingFullAddress"`
	CodAmount             float64 `json:"codAmount"`
	ExpectedDeliveryTime  string  `json:"expectedDeliveryTime"`
	CreatedDate           string  `json:"createdDate"`
	ErrorMessage          string  `json:"errorMessage"`
}
