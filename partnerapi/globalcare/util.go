package globalcare

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/Selly-Modules/logger"
)

// toBytes ...
func toBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		logger.Error("pjson.toBytes", logger.LogData{"payload": data})
	}
	return b
}

// toJSONString ...
func toJSONString(data interface{}) string {
	return string(toBytes(data))
}

// GeneratePublicKeyFromBytes ...
func generatePublicKeyFromBytes(b []byte) (*rsa.PublicKey, error) {
	pubPem, _ := pem.Decode(b)
	p, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("prsa.GeneratePublicKeyFromBytes: ParsePKIXPublicKey %v", err)
	}
	return p.(*rsa.PublicKey), nil
}

// GeneratePrivateKeyFromBytes ...
func generatePrivateKeyFromBytes(b []byte) (*rsa.PrivateKey, error) {
	privPem, _ := pem.Decode(b)
	p, err := x509.ParsePKCS1PrivateKey(privPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("prsa.GeneratePrivateKeyFromBytes: ParsePKCS1PrivateKey %v", err)
	}
	return p, nil
}
