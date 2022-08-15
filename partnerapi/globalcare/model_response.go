package globalcare

// CommonResponse ...
type CommonResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
