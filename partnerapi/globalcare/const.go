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
	VehicleTypeID50ccAbove    = 1
	VehicleTypeID50ccBelow    = 2
	VehicleTypeIDElectricBike = 3
)

const (
	CarOccupantAccidentInsurance0   = 1
	CarOccupantAccidentInsurance10m = 2
	CarOccupantAccidentInsurance20m = 3
)

const (
	productCodeDefault = "bbxm"
	providerIDDefault  = 4
	productIDDefault   = 18
)
