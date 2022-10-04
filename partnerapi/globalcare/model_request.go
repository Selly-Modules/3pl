package globalcare

import (
	"encoding/json"
	"time"
)

// CommonRequestBody ...
type CommonRequestBody struct {
	Signature string `json:"signature"`
	Data      string `json:"data"`
}

// CreateOrderPayload ...
type CreateOrderPayload struct {
	ProductCode string      `json:"productCode"`
	ProviderID  int         `json:"providerId"`
	ProductID   int         `json:"productId"`
	PartnerID   string      `json:"partnerId"`
	VehicleInfo VehicleInfo `json:"vehicleInfo"`
	InsuredInfo InsuredInfo `json:"insuredInfo"`
}

// VehicleInfo ...
type VehicleInfo struct {
	TypeID   int    `json:"typeId"`
	TypeName string `json:"typeName"`
	License  string `json:"license"`
	Chassis  string `json:"chassis"`
	Engine   string `json:"engine"`

	// V2 = true if TypeID = 1 and insurance type is car
	V2 bool `json:"v2,omitempty"`
	// CarOccupantAccidentInsurance type int for motorbike, type CarOccupantAccidentInsuranceObj for car insurance
	CarOccupantAccidentInsurance interface{} `json:"carOccupantAccidentInsurance"`
	NumberOfSeatsOver25          int         `json:"numberOfSeatsOver25"`
	NumberOfSeatsOrTonnageName   string      `json:"numberOfSeatsOrTonnageName"`
	NumberOfSeatsOrTonnage       int         `json:"numberOfSeatsOrTonnage"`
}

// CarOccupantAccidentInsuranceObj ...
type CarOccupantAccidentInsuranceObj struct {
	NumberOfSeats int `json:"numberOfSeats"`
}

func (c *CarOccupantAccidentInsuranceObj) MarshalJSON() ([]byte, error) {
	buy := 1
	if c.NumberOfSeats <= 0 {
		buy = 2
	}
	return json.Marshal(&struct {
		Buy           int `json:"buy"`
		NumberOfSeats int `json:"numberOfSeats"`
	}{
		Buy:           buy,
		NumberOfSeats: c.NumberOfSeats,
	})
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
	Status           int       `json:"status"`
	OrderCode        string    `json:"orderCode"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Note             string    `json:"note"`
	CertLink         string    `json:"certLink"`
	PartnerOrderCode string    `json:"partnerOrderCode"`
}
