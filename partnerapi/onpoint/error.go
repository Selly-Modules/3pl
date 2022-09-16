package onpoint

import "fmt"

// Error ...
type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Error ...
func (e Error) Error() string {
	return fmt.Sprintf("onpoint_err: status %s, message: %s", e.Status, e.Message)
}
