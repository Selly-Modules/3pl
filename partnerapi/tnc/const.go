package tnc

const (
	TimeLayout = "2006-01-02 15:04:05"

	apiPathCreateOutboundRequest = "/api/v1/ors"
	apiPathGetOutboundRequest    = "/api/v1/ors/%d"
	apiPathCancelOutboundRequest = "/api/v1/ors/%d/cancel"
	apiPathAuth                  = "/auth/realms/%s/protocol/openid-connect/token"

	PriorityUrgent = 3
	PriorityHigh   = 2
	PriorityNormal = 1

	TPLCodeGHN          = "GHN"
	TPLCodeGHTK         = "GHTK"
	TPLCodeBest         = "BEST"
	TPLCodeSnappy       = "SPY"
	TPLCodeViettelPost  = "VTP"
	TPLCodeSellyExpress = "SE"
	TPLCodeJTExpress    = "JTE"
)

const (
	baseURLAuthStaging = "https://auth.stg.tnclog.vn"
	baseURLStaging     = "https://ext-api.stg.tnclog.vn"

	baseURLAuthProd = "https://auth.tnclog.vn"
	baseURLProd     = "https://ext-api.tnclog.vn"
)

var (
	baseURLENVMapping = map[ENV]string{
		EnvProd:    baseURLProd,
		EnvStaging: baseURLStaging,
	}
	baseURLAuthENVMapping = map[ENV]string{
		EnvProd:    baseURLAuthProd,
		EnvStaging: baseURLAuthStaging,
	}
)
