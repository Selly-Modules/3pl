package onpoint

import (
	"fmt"
	"strings"
)

// Error ...
type Error struct {
	Message string              `json:"message"`
	Code    string              `json:"code"`
	Errors  map[string][]string `json:"errors"`
}

// Error ...
func (e Error) Error() string {
	msg := fmt.Sprintf("onpoint_err: code %s, message: %s", e.Code, e.Message)
	if len(e.Errors) > 0 {
		msg += "\ndetail: "
		for k, v := range e.Errors {
			msg += fmt.Sprintf("field %s - error %s", k, strings.Join(v, ","))
		}
	}
	return msg
}
