package natsiomodel

import "github.com/Selly-Modules/tpl/util/pjson"

// NatsResponse ...
type NatsResponse struct {
	Data      interface{} `json:"data"`
	Error     bool        `json:"error"`
	Message   string      `json:"message"`
	RequestID string      `json:"requestId"`
}

// ParseData ...
func (r *NatsResponse) ParseData(result interface{}) error {
	b := pjson.ToBytes(r.Data)
	return pjson.Unmarshal(b, result)
}
