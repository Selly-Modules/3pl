package natsiomodel

import "github.com/Selly-Modules/tpl/util/pjson"

// NatsResponse ...
type NatsResponse struct {
	Response  *HttpResponse `json:"response"`
	Error     bool          `json:"error"`
	Message   string        `json:"message"`
	RequestID string        `json:"requestId"`
}

// ParseResponseData ...
func (r *NatsResponse) ParseResponseData(result interface{}) error {
	if r.Response == nil {
		return nil
	}
	b := pjson.ToBytes(r.Response.Body)
	return pjson.Unmarshal(b, result)
}

// HttpResponse ...
type HttpResponse struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
}
