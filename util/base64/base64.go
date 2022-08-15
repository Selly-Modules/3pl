package base64

import b64 "encoding/base64"

// Encode ...
func Encode(data []byte) string {
	sEnc := b64.StdEncoding.EncodeToString(data)
	return sEnc
}

// Decode ...
func Decode(text string) []byte {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	return sDec
}
