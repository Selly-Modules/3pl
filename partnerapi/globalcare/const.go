package globalcare

const (
	baseURLDev     = "https://core-dev.globalcare.vn"
	baseURLStaging = "https://core-beta.globalcare.vn"
	baseURLProd    = "https://core.globalcare.vn"
)

var (
	baseURLENVMapping = map[ENV]string{
		EnvProd:    baseURLProd,
		EnvStaging: baseURLStaging,
		EnvDev:     baseURLDev,
	}
)

const (
	apiPathCreateOrder = "/api/v1/partner/selly/order/create"
	apiPathGetOrder    = "/api/v1/partner/selly/order/%s/detail"
)

const (
	MotorbikeProductCode = "bbxm"
	MotorbikeProviderID  = 4
	MotorbikeProductID   = 18

	CarProductCode = "bbot"
	CarProviderID  = 10
	CarProductID   = 35

	NumOfSeatsMinValue = 25
)

type InsuranceConfig struct {
	ProductCode string
	ProviderID  int
	ProductID   int
}

var (
	CarConfig = InsuranceConfig{
		ProductCode: CarProductCode,
		ProviderID:  CarProviderID,
		ProductID:   CarProductID,
	}

	MotorbikeConfig = InsuranceConfig{
		ProductCode: MotorbikeProductCode,
		ProviderID:  MotorbikeProviderID,
		ProductID:   MotorbikeProductID,
	}
)
