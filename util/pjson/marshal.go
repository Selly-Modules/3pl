package pjson

import (
	"encoding/json"
	"log"
)

// Marshal ...
func Marshal(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("3pl/util/pjson.Marshal: err %v, payload %v", err, data)
	}
	return b, err
}
