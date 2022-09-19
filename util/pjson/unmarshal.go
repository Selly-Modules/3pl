package pjson

import (
	"encoding/json"
	"log"
)

// Unmarshal ...
func Unmarshal(b []byte, resultP interface{}) error {
	err := json.Unmarshal(b, resultP)
	if err != nil {
		log.Printf("3pl/util/pjson.Unmarshal: err %v, payload %s", err, string(b))
	}
	return err
}
