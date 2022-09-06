package globalcare

import "time"

// CommonRequestBody ...
type CommonRequestBody struct {
	Signature string `json:"signature"`
	Data      string `json:"data"`
}

// CreateOrderPayload ...
type CreateOrderPayload struct {
	PartnerOrderCode string      `json:"partnerOrderCode"`
	VehicleInfo      VehicleInfo `json:"vehicleInfo"`
	InsuredInfo      InsuredInfo `json:"insuredInfo"`
}

type createOrderData struct {
	ProductCode string      `json:"productCode"`
	ProviderID  int         `json:"providerId"`
	ProductID   int         `json:"productId"`
	PartnerID   string      `json:"partnerId"`
	VehicleInfo VehicleInfo `json:"vehicleInfo"`
	InsuredInfo InsuredInfo `json:"insuredInfo"`
}

// VehicleInfo ...
type VehicleInfo struct {
	TypeID                       int    `json:"typeId"`
	TypeName                     string `json:"typeName"`
	CarOccupantAccidentInsurance int    `json:"carOccupantAccidentInsurance"`
	License                      string `json:"license"`
	Chassis                      string `json:"chassis"`
	Engine                       string `json:"engine"`
}

// InsuredInfo ...
type InsuredInfo struct {
	BuyerName        string `json:"buyerName"`
	BuyerPhone       string `json:"buyerPhone"`
	BuyerEmail       string `json:"buyerEmail"`
	BuyerAddress     string `json:"buyerAddress"`
	YearsOfInsurance string `json:"yearsOfInsurance"`
	BeginDate        string `json:"beginDate"`
}

// Webhook ...
type Webhook struct {
	Status    int       `json:"status"`
	OrderCode string    `json:"orderCode"`
	UpdatedAt time.Time `json:"updatedAt"`
	Note      string    `json:"note"`
	CertLink  string    `json:"certLink"`
}
