package globalcare

import (
	"encoding/json"

	"github.com/Selly-Modules/3pl/util/base64"
)

// CommonResponse ...
type CommonResponse struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

// DecodeCreateOrderSuccess ...
func (r *CommonResponse) DecodeCreateOrderSuccess() (res CreateOrderResponseDecoded, err error) {
	err = r.Decode(&res)
	return res, err
}

// DecodeGetOrderSuccess ...
func (r *CommonResponse) DecodeGetOrderSuccess() (res GetOrderResponseDecoded, err error) {
	err = r.Decode(&res)
	return res, err
}

// DecodeError ...
func (r *CommonResponse) DecodeError() (res ResponseError, err error) {
	err = r.Decode(&res)
	return res, err
}

// Decode ...
func (r *CommonResponse) Decode(resultPointer interface{}) error {
	b := base64.Decode(r.Data)
	return json.Unmarshal(b, resultPointer)
}

// CreateOrderResponseDecoded ...
type CreateOrderResponseDecoded struct {
	StatusCode int               `json:"statusCode"`
	Result     CreateOrderResult `json:"result"`
}

// CreateOrderResult ...
type CreateOrderResult struct {
	OrderCode   string `json:"orderCode"`
	PaymentLink string `json:"paymentLink"`
	Fees        int    `json:"fees"`
	StatusId    int    `json:"statusId"`
}

// ResponseError ...
type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// GetOrderResponseDecoded ...
type GetOrderResponseDecoded struct {
	StatusCode int            `json:"statusCode"`
	Result     GetOrderResult `json:"result"`
}

// GetOrderResult ...
type GetOrderResult struct {
	ProviderTitle string        `json:"providerTitle"`
	BeginDate     string        `json:"beginDate"`
	EndDate       string        `json:"endDate"`
	Amount        string        `json:"amount"`
	CertLink      string        `json:"certLink"`
	StatusId      int           `json:"statusId"`
	StatusTitle   string        `json:"statusTitle"`
	Buyer         BuyerInfo     `json:"buyer"`
	InsuredInfo   InsuranceInfo `json:"insuredInfo"`
}

// InsuranceInfo ...
type InsuranceInfo struct {
	TypeId   int    `json:"typeId"`
	TypeName string `json:"typeName"`
}

// BuyerInfo ...
type BuyerInfo struct {
	BuyerName      string      `json:"buyerName"`
	BuyerPrivateId interface{} `json:"buyerPrivateId"`
	BuyerPhone     string      `json:"buyerPhone"`
	BuyerAddress   string      `json:"buyerAddress"`
	BuyerEmail     string      `json:"buyerEmail"`
}
