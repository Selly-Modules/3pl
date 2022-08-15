package globalcare

import (
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/natsio"
	"github.com/nats-io/nats.go"
	"github.com/thoas/go-funk"

	"github.com/Selly-Modules/tpl/constant"
	natsiomodel "github.com/Selly-Modules/tpl/model/natsio"
	"github.com/Selly-Modules/tpl/util/base64"
	"github.com/Selly-Modules/tpl/util/pjson"
)

// Client ...
type Client struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	env        ENV
	natsClient natsio.Server
}

// NewClient generate Client
// using privateKey to decrypt data from Global Care
// using publicKey to encrypt data before send to Global Care
func NewClient(env ENV, privateKey, publicKey string) (*Client, error) {
	validENVs := []ENV{EnvProd, EnvDev, EnvStaging}
	if !funk.Contains(validENVs, env) {
		return nil, fmt.Errorf("globalcare.NewClient - invalid_env: %s", env)
	}
	pubKey, err := generatePublicKeyFromBytes([]byte(publicKey))
	if err != nil {
		return nil, fmt.Errorf("globalcare.NewClient - invalid_public_key: %v", err)
	}
	privKey, err := generatePrivateKeyFromBytes([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("globalcare.NewClient - invalid_private_key: %v", err)
	}
	return &Client{
		privateKey: privKey,
		publicKey:  pubKey,
	}, nil
}

// CreateOrder ...
func (c *Client) CreateOrder(p CreateOrderPayload) (*CommonResponse, error) {
	url := c.getBaseURL() + apiPathCreateOrder
	data := createOrderData{
		ProductCode: productCodeDefault,
		ProviderID:  providerIDDefault,
		ProductID:   productIDDefault,
		PartnerID:   p.PartnerOrderCode,
		VehicleInfo: p.VehicleInfo,
		InsuredInfo: p.InsuredInfo,
	}

	body := CommonRequestBody{
		Signature: "", // TODO:implement
		Data:      base64.Encode(pjson.ToBytes(data)),
	}
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    url,
			Method: http.MethodPost,
			Data:   pjson.ToJSONString(body),
		},
	}
	msg, err := c.requestNats(constant.NatsCommunicationSubjectRequestHTTP, natsPayload)
	if err != nil {
		logger.Error("globalcare.Client.CreateOrder", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r   natsiomodel.NatsResponse
		res CommonResponse
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	err = r.ParseResponseData(&res)
	return &res, err
}

// GetOrder ...
func (c *Client) GetOrder(orderCode string) (*CommonResponse, error) {
	url := c.getBaseURL() + fmt.Sprintf(apiPathGetOrder, orderCode)
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    url,
			Method: http.MethodGet,
		},
	}
	msg, err := c.requestNats(constant.NatsCommunicationSubjectRequestHTTP, natsPayload)
	if err != nil {
		logger.Error("globalcare.Client.GetOrder", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r   natsiomodel.NatsResponse
		res CommonResponse
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	err = r.ParseResponseData(&res)
	return &res, err
}

func (c *Client) requestNats(subject string, data interface{}) (*nats.Msg, error) {
	b := toBytes(data)
	return c.natsClient.Request(subject, b)
}

func (c *Client) getBaseURL() string {
	return baseURLENVMapping[c.env]
}
