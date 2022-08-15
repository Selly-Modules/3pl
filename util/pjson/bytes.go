package pjson

import (
	"encoding/json"

	"github.com/Selly-Modules/logger"
)

// ToBytes ...
func ToBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		logger.Error("pjson.ToBytes", logger.LogData{"payload": data})
	}
	return b
}
