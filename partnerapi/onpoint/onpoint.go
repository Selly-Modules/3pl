package onpoint

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/Selly-Modules/natsio/subject"

	"github.com/Selly-Modules/3pl/util/pjson"
)

// Client ...
type Client struct {
	env        ENV
	apiKey     string
	secretKey  string
	natsClient natsio.Server
}

// NewClient generate OnPoint client
func NewClient(env ENV, apiKey, secretKey string, nc natsio.Server) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("onpoint: cannot init with empty api key")
	}
	return &Client{
		env:        env,
		apiKey:     apiKey,
		secretKey:  secretKey,
		natsClient: nc,
	}, nil
}

// CreateOrder ...
func (c *Client) CreateOrder(p CreateOrderRequest) (*CreateOrderResponse, error) {
	url := c.getBaseURL() + apiPathCreateOrder
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodPost,
			Data:   pjson.ToJSONString(p),
		},
	}
	var (
		r       model.CommunicationHttpResponse
		errRes  Error
		dataRes struct {
			Data CreateOrderResponse `json:"data"`
		}
	)
	if err := c.requestHttpViaNats(natsPayload, &r); err != nil {
		return nil, err
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("onpoint.Client.CreateOrder: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err := r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("onpoint.Client.CreateOrder: parse_response_err: %v", err)
		}
		return nil, errRes
	}
	if err := r.ParseResponseData(&dataRes); err != nil {
		return nil, fmt.Errorf("onpoint.Client.CreateOrder: parse_response_data: %v", err)
	}

	return &dataRes.Data, nil
}

// UpdateDelivery ...
func (c *Client) UpdateDelivery(p UpdateOrderDeliveryRequest) (*UpdateOrderDeliveryResponse, error) {
	url := c.getBaseURL() + apiPathUpdateDelivery
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodPost,
			Data:   pjson.ToJSONString(p),
		},
	}
	var (
		r       model.CommunicationHttpResponse
		errRes  Error
		dataRes struct {
			Data UpdateOrderDeliveryResponse `json:"data"`
		}
	)
	if err := c.requestHttpViaNats(natsPayload, &r); err != nil {
		return nil, err
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("onpoint.Client.UpdateDelivery: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err := r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("onpoint.Client.UpdateDelivery: parse_response_err: %v", err)
		}
		return nil, errRes
	}
	if err := r.ParseResponseData(&dataRes); err != nil {
		return nil, fmt.Errorf("onpoint.Client.UpdateDelivery: parse_response_data: %v", err)
	}

	return &dataRes.Data, nil
}

// CancelOrder ...
func (c *Client) CancelOrder(p CancelOrderRequest) (*CancelOrderResponse, error) {
	url := c.getBaseURL() + apiPathCancelOrder
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodPost,
			Data:   pjson.ToJSONString(p),
		},
	}
	var (
		r       model.CommunicationHttpResponse
		errRes  Error
		dataRes struct {
			Data CancelOrderResponse `json:"data"`
		}
	)
	if err := c.requestHttpViaNats(natsPayload, &r); err != nil {
		return nil, err
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("onpoint.Client.CancelOrder: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err := r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("onpoint.Client.CancelOrder: parse_response_err: %v", err)
		}
		return nil, errRes
	}
	if err := r.ParseResponseData(&dataRes); err != nil {
		return nil, fmt.Errorf("onpoint.Client.CancelOrder: parse_response_data: %v", err)
	}

	return &dataRes.Data, nil
}

// GetChannels ...
func (c *Client) GetChannels() ([]ChannelResponse, error) {
	url := c.getBaseURL() + apiPathGetChannels
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    url,
			Method: http.MethodGet,
		},
	}
	var (
		r       model.CommunicationHttpResponse
		errRes  Error
		dataRes struct {
			Data []ChannelResponse `json:"data"`
		}
	)
	if err := c.requestHttpViaNats(natsPayload, &r); err != nil {
		return nil, err
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("onpoint.Client.GetChannels: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err := r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("onpoint.Client.GetChannels: parse_response_err: %v", err)
		}
		return nil, errRes
	}
	if err := r.ParseResponseData(&dataRes); err != nil {
		return nil, fmt.Errorf("onpoint.Client.GetChannels: parse_response_data: %v", err)
	}

	return dataRes.Data, nil
}

func (c *Client) requestHttpViaNats(data model.CommunicationRequestHttp, res interface{}) error {
	ec, err := c.natsClient.NewJSONEncodedConn()
	if err != nil {
		return fmt.Errorf("onpoint: request via nats %v", err)
	}
	qs := ""
	for k, v := range data.Payload.Query {
		qs += k + "=" + v
	}
	now := time.Now().Unix()
	ts := strconv.FormatInt(now, 10)
	arr := []string{
		qs,
		data.Payload.Data,
		ts,
	}
	s := strings.Join(arr, ".")
	// sign data
	sign := hashSHA256AndUppercase(s, c.secretKey)
	data.Payload.Header = map[string]string{
		headerXAPIKey:    c.apiKey,
		headerXSignature: sign,
		headerXTimestamp: ts,
	}

	return ec.Request(subject.Communication.RequestHTTP, data, res)
}

func (c *Client) getBaseURL() string {
	return baseURLENVMapping[c.env]
}
