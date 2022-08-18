package tnc

// Product ...
type Product struct {
	PartnerSKU        string `json:"partnerSKU"`
	UnitCode          string `json:"unitCode"`
	ConditionTypeCode string `json:"conditionTypeCode"`
	Quantity          int64  `json:"quantity"`
}

// Address ...
type Address struct {
	AddressNo    string `json:"addressNo"`
	ProvinceCode string `json:"provinceCode"`
	DistrictCode string `json:"districtCode"`
	WardCode     string `json:"wardCode"`
}

// OutboundRequestPayload ...
type OutboundRequestPayload struct {
	WarehouseCode       string    `json:"warehouseCode"`
	ShippingServiceCode string    `json:"shippingServiceCode"`
	PartnerORCode       string    `json:"partnerORCode"`
	PartnerRefId        string    `json:"partnerRefId"`
	RefCode             string    `json:"refCode"`
	CodAmount           float64   `json:"codAmount"`
	PriorityType        int       `json:"priorityType"`
	CustomerName        string    `json:"customerName"`
	CustomerPhoneNumber string    `json:"customerPhoneNumber"`
	Type                int       `json:"type"`
	ShippingType        int       `json:"shippingType"`
	VehicleNumber       string    `json:"vehicleNumber"`
	ContainerNumber     string    `json:"containerNumber"`
	PackType            int       `json:"packType"`
	PackingNote         string    `json:"packingNote"`
	CustomLabel         bool      `json:"customLabel"`
	BizType             int       `json:"bizType"`
	Note                string    `json:"note"`
	ShippingAddress     Address   `json:"shippingAddress"`
	Products            []Product `json:"products"`
	PartnerCreationTime string    `json:"partnerCreationTime"`
	TPLCode             string    `json:"tplCode"`
	TrackingCode        string    `json:"trackingCode"`
}

// UpdateORLogisticInfoPayload ...
type UpdateORLogisticInfoPayload struct {
	OrID          int    `json:"orId"`
	TrackingCode  string `json:"trackingCode"`
	ShippingLabel string `json:"shippingLabel"`
	SlaShipDate   string `json:"slaShipDate"`
}
