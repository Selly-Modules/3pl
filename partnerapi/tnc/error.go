package tnc

import (
	"fmt"
)

// Error ...
type Error struct {
	Code    string `json:"code"`
	Message string `json:"errorMessage"`
}

// Error ...
func (e Error) Error() string {
	return fmt.Sprintf("tnc_err: code %s, messsage %s", e.Code, e.Message)
}

// IsErrExistPartnerCode ...
func IsErrExistPartnerCode(err error) bool {
	e, ok := err.(Error)
	return ok && e.Code == ErrCodeExistPartnerCode
}
