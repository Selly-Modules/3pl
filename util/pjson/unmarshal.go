package pjson

import (
	"encoding/json"

	"github.com/Selly-Modules/logger"
)

// Unmarshal ...
func Unmarshal(b []byte, resultP interface{}) error {
	err := json.Unmarshal(b, resultP)
	if err != nil {
		logger.Error("pjson.Unmarshal", logger.LogData{
			"raw": string(b),
			"err": err.Error(),
		})
	}
	return err
}
