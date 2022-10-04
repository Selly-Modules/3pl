package onpoint

const (
	baseURLStaging = "https://dev-selly-api.onpoint.vn"
	baseURLProd    = "https://selly-api.onpoint.vn"

	apiPathCreateOrder    = "/v1/orders/create"
	apiPathUpdateDelivery = "/v1/orders/update_delivery"
	apiPathCancelOrder    = "/v1/orders/cancel"
	apiPathGetChannels    = "/v1/channels"

	headerXAPIKey    = "x-api-key"
	headerXTimestamp = "x-timestamp"
	headerXSignature = "x-signature"

	webhookEventUpdateOrderStatus = "update_order_status"
	webhookEventUpdateInventory   = "update_inventory"

	CodeSuccess = "SUCCESS"
)

var (
	baseURLENVMapping = map[ENV]string{
		EnvProd:    baseURLProd,
		EnvStaging: baseURLStaging,
	}
)
