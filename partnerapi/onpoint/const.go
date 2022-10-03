package onpoint

const (
	baseURLStaging = "https://dev-selly-api.onpoint.vn"
	baseURLProd    = "https://selly-api.onpoint.vn"

	apiPathCreateOrder    = "/v1/orders"
	apiPathUpdateDelivery = "/v1/orders/delivery/update"
	apiPathCancelOrder    = "/v1/orders/cancel"
	apiPathGetChannels    = "/v1/channels"

	headerXAPIKey    = "x-api-key"
	headerXTimestamp = "x-timestamp"
	headerXSignature = "x-signature"

	webhookEventUpdateOrderStatus = "update_order_status"
	webhookEventUpdateInventory   = "update_inventory"
)

var (
	baseURLENVMapping = map[ENV]string{
		EnvProd:    baseURLProd,
		EnvStaging: baseURLStaging,
	}
)
