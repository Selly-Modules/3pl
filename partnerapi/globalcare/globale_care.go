package globalcare

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"
	"github.com/nats-io/nats.go"
	"github.com/thoas/go-funk"

	"github.com/Selly-Modules/3pl/util/base64"
	"github.com/Selly-Modules/3pl/util/pjson"
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
func NewClient(env ENV, privateKey, publicKey string, natsClient natsio.Server) (*Client, error) {
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
		env:        env,
		natsClient: natsClient,
	}, nil
}

// CreateOrder ...
func (c *Client) CreateOrder(p CreateOrderPayload) (*CreateOrderResponseDecoded, error) {
	url := c.getBaseURL() + apiPathCreateOrder
	data := createOrderData{
		ProductCode: productCodeDefault,
		ProviderID:  providerIDDefault,
		ProductID:   productIDDefault,
		PartnerID:   p.PartnerOrderCode,
		VehicleInfo: p.VehicleInfo,
		InsuredInfo: p.InsuredInfo,
	}

	dataString := base64.Encode(pjson.ToBytes(data))
	sign, err := c.signData(dataString)
	if err != nil {
		return nil, fmt.Errorf("globalcare.Client.CreateOrder - sign_err %v", err)
	}
	body := CommonRequestBody{
		Signature: sign,
		Data:      dataString,
	}
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodPost,
			Data:   pjson.ToJSONString(body),
			Header: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	msg, err := c.requestNats(subject.Communication.RequestHTTP, natsPayload)
	if err != nil {
		logger.Error("globalcare.Client.CreateOrder", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r   model.CommunicationHttpResponse
		res CommonResponse
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		log.Printf("globalcare.Client.CreateOrder - pjson.Unmarshal %v, %s\n", err, string(msg.Data))
		return nil, err
	}
	if err = r.ParseResponseData(&res); err != nil {
		log.Printf("globalcare.Client.CreateOrder - ParseResponseData %v, %s\n", err, string(msg.Data))
		return nil, err
	}
	if r.Response == nil {
		log.Println("globalcare.Client.CreateOrder - nil response")
		return nil, fmt.Errorf("globalcare.Client.CreateOrder create_order_empty_response")
	}

	if r.Response.StatusCode >= http.StatusBadRequest {
		log.Println("globalcare.Client.CreateOrder - bad request", res)
		info, err := res.DecodeError()
		if err != nil {
			log.Println("globalcare.Client.CreateOrder - decode err", err)
			return nil, err
		}
		return nil, errors.New(info.Message)
	}
	info, err := res.DecodeCreateOrderSuccess()
	if err != nil {
		log.Println("globalcare.Client.CreateOrder - DecodeCreateOrderSuccess err:", err, string(msg.Data))
		return nil, err
	}

	return &info, err
}

// GetOrder ...
func (c *Client) GetOrder(orderCode string) (*GetOrderResponseDecoded, error) {
	url := c.getBaseURL() + fmt.Sprintf(apiPathGetOrder, orderCode)
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodGet,
		},
	}
	msg, err := c.requestNats(subject.Communication.RequestHTTP, natsPayload)
	if err != nil {
		logger.Error("globalcare.Client.GetOrder", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r   model.CommunicationHttpResponse
		res CommonResponse
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, err
	}
	if err = r.ParseResponseData(&res); err != nil {
		return nil, err
	}
	if r.Response == nil {
		return nil, fmt.Errorf("globalcare.Client.GetOrder get_order_empty_response")
	}

	if r.Response.StatusCode >= http.StatusBadRequest {
		info, err := res.DecodeError()
		if err != nil {
			return nil, err
		}
		return nil, errors.New(info.Message)
	}
	info, err := res.DecodeGetOrderSuccess()
	if err != nil {
		return nil, err
	}
	return &info, err
}

func (c *Client) requestNats(subject string, data interface{}) (*nats.Msg, error) {
	b := toBytes(data)
	return c.natsClient.Request(subject, b)
}

func (c *Client) getBaseURL() string {
	return baseURLENVMapping[c.env]
}

func (c *Client) signData(s string) (string, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(s))
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, msgHashSum)
	if err != nil {
		return "", err
	}

	return base64.Encode(signature), nil
}
