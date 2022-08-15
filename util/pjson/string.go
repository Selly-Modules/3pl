package pjson

// ToJSONString ...
func ToJSONString(data interface{}) string {
	return string(ToBytes(data))
}
