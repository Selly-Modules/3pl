package natsiomodel

// NatsRequestHTTP ...
type NatsRequestHTTP struct {
	ResponseImmediately bool        `json:"responseImmediately"`
	AuthenticationID    string      `json:"authenticationId"`
	Payload             HTTPPayload `json:"payload"`
}

// HTTPPayload ...
type HTTPPayload struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Data   string            `json:"data"`
	Header map[string]string `json:"header"`
	Query  map[string]string `json:"query"`
}
