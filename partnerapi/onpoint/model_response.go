package onpoint

// CreateOrderResponse ...
type CreateOrderResponse struct {
	OrderCode        string      `json:"order_code"`
	OnpointOrderCode string      `json:"onpoint_order_code"`
	OrderDate        string      `json:"order_date"`
	Note             string      `json:"note"`
	SubtotalPrice    int         `json:"subtotal_price"`
	TotalDiscounts   int         `json:"total_discounts"`
	TotalPrice       int         `json:"total_price"`
	PaymentMethod    string      `json:"payment_method"`
	DeliveryPlatform string      `json:"delivery_platform"`
	Status           string      `json:"status"`
	UpdatedAt        string      `json:"updated_at"`
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
