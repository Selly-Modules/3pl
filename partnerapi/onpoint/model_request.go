package onpoint

import (
	"time"

	"github.com/Selly-Modules/3pl/util/pjson"
)

/*
 * Request payload
 */

// CreateOrderRequest ...
type CreateOrderRequest struct {
	PartnerOrderCode string      `json:"partner_order_code"`
	OrderDate        time.Time   `json:"order_date"`
	ChannelCode      string      `json:"channel_code"`
	FullName         string      `json:"full_name"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	Address          string      `json:"address"`
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
	Items            []OrderItem `json:"items"`
}

// OrderItem ...
type OrderItem struct {
	SellingPrice int    `json:"selling_price"`
	Quantity     int    `json:"quantity"`
	Uom          string `json:"uom"`
	Amount       int    `json:"amount"`
	Name         string `json:"name"`
	PartnerSku   string `json:"partner_sku"`
}

// UpdateOrderDeliveryRequest ...
type UpdateOrderDeliveryRequest struct {
	OrderNo                string `json:"order_no"`          // required
	DeliveryPlatform       string `json:"delivery_platform"` // required
	DeliveryTrackingNumber string `json:"delivery_tracking_number"`
	ShippingLabel          string `json:"shipping_label"`
}

// CancelOrderRequest ...
type CancelOrderRequest struct {
	OrderNo string `json:"order_no"`
}

/**
 * WEBHOOK ONPOINT
 */

// WebhookDataUpdateInventory ...
type WebhookDataUpdateInventory struct {
	Sku               string    `json:"sku"`
	PartnerSku        string    `json:"partner_sku"`
	WarehouseCode     string    `json:"warehouse_code"`
	AvailableQuantity int       `json:"available_quantity"`
	CommittedQuantity int       `json:"committed_quantity"`
	TotalQuantity     int       `json:"total_quantity"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// WebhookDataUpdateOrderStatus ...
type WebhookDataUpdateOrderStatus struct {
	PartnerOrderCode string    `json:"partner_order_code"`
	OrderNo          string    `json:"order_no"`
	Status           string    `json:"status"`
	DeliveryStatus   string    `json:"delivery_status"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// WebhookPayload ...
type WebhookPayload struct {
	Event       string      `json:"event"`
	RequestedAt time.Time   `json:"requested_at"`
	Data        interface{} `json:"data"`
}

// GetDataEventUpdateOrderStatus ...
func (p WebhookPayload) GetDataEventUpdateOrderStatus() (data *WebhookDataUpdateOrderStatus, ok bool) {
	if p.Event != webhookEventUpdateOrderStatus {
		return nil, false
	}
	b, err := pjson.Marshal(p.Data)
	if err != nil {
		return nil, false
	}
	if err = pjson.Unmarshal(b, &data); err != nil {
		return nil, false
	}
	return data, true
}

// GetDataEventUpdateInventory ...
func (p WebhookPayload) GetDataEventUpdateInventory() (data *WebhookDataUpdateInventory, ok bool) {
	if p.Event != webhookEventUpdateInventory {
		return nil, false
	}
	b, err := pjson.Marshal(p.Data)
	if err != nil {
		return nil, false
	}
	if err = pjson.Unmarshal(b, &data); err != nil {
		return nil, false
	}
	return data, true
}
