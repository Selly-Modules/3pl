package onpoint

import "time"

// CreateOrderResponse ...
type CreateOrderResponse struct {
	PartnerOrderCode string      `json:"partner_order_code"`
	OrderNo          string      `json:"order_no"`
	OrderDate        time.Time   `json:"order_date"`
	ChannelCode      string      `json:"channel_code"`
	FullName         string      `json:"full_name"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	Address          string      `json:"address"`
	FullAddress      string      `json:"full_address"`
	District         string      `json:"district"`
	Ward             string      `json:"ward"`
	Province         string      `json:"province"`
	DistrictCode     string      `json:"district_code"`
	WardCode         string      `json:"ward_code"`
	ProvinceCode     string      `json:"province_code"`
	Note             string      `json:"note"`
	SubtotalPrice    int         `json:"subtotal_price"`
	ShippingFee      int         `json:"shipping_fee"`
	TotalDiscounts   int         `json:"total_discounts"`
	TotalPrice       int         `json:"total_price"`
	PaymentMethod    string      `json:"payment_method"`
	DeliveryPlatform string      `json:"delivery_platform"`
	Status           string      `json:"status"`
	UpdatedAt        time.Time   `json:"updated_at"`
	InsertedAt       time.Time   `json:"inserted_at"`
	Items            []OrderItem `json:"items"`
}

// UpdateOrderDeliveryResponse ...
type UpdateOrderDeliveryResponse struct {
	DeliveryPlatform       string `json:"delivery_platform"`
	DeliveryTrackingNumber string `json:"delivery_tracking_number"`
	ShippingLabel          string `json:"shipping_label"`
}

// CancelOrderResponse ...
type CancelOrderResponse struct {
	OrderNo string `json:"order_no"`
	Status  string `json:"status"`
}

// ChannelResponse ...
type ChannelResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
