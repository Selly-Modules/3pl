package tnc

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/natsio"
	"github.com/nats-io/nats.go"

	"github.com/Selly-Modules/tpl/constant"
	natsiomodel "github.com/Selly-Modules/tpl/model/natsio"
	"github.com/Selly-Modules/tpl/util/pjson"
)

// Client ...
type Client struct {
	// Auth info
	realm        string
	clientID     string
	clientSecret string

	env        ENV
	natsClient natsio.Server

	token         string
	tokenExpireAt time.Time
}

// NewClient ...
func NewClient(env ENV, clientID, clientSecret, realm string, natsClient natsio.Server) (*Client, error) {
	if env != EnvProd && env != EnvStaging {
		return nil, fmt.Errorf("tnc.NewClient: invalid_env %s", env)
	}
	return &Client{
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
		token:        "",
		env:          env,
		natsClient:   natsClient,
	}, nil
}

// CreateOutboundRequest ...
func (c *Client) CreateOutboundRequest(p OutboundRequestPayload) (*OutboundRequestRes, error) {
	apiURL := c.getBaseURL() + apiPathCreateOutboundRequest
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    apiURL,
			Method: http.MethodGet,
			Data:   pjson.ToJSONString(p),
			Header: c.getRequestHeader(),
		},
	}
	msg, err := c.requestHttpViaNats(natsPayload)
	if err != nil {
		logger.Error("tnc.Client.CreateOutboundRequest - requestHttpViaNats", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r       natsiomodel.NatsResponse
		errRes  ErrRes
		dataRes []OutboundRequestRes
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: parse_data %v", err)
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err = r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: parse_response_err: %v", err)
		}
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: failed code %s, message %s", errRes.Code, errRes.ErrorMessage)
	}
	if err = r.ParseResponseData(&dataRes); err != nil {
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: parse_response_data: %v", err)
	}
	if len(dataRes) == 0 {
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: empty_result")
	}

	return &dataRes[0], err
}

// GetOutboundRequestByID ...
func (c *Client) GetOutboundRequestByID(requestID int) (*OutboundRequestInfo, error) {
	apiURL := c.getBaseURL() + fmt.Sprintf(apiPathGetOutboundRequest, requestID)
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    apiURL,
			Method: http.MethodGet,
			Header: c.getRequestHeader(),
		},
	}
	msg, err := c.requestHttpViaNats(natsPayload)
	if err != nil {
		logger.Error("tnc.Client.GetOutboundRequestByID - requestHttpViaNats", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r               natsiomodel.NatsResponse
		errRes          ErrRes
		outboundRequest OutboundRequestInfo
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, fmt.Errorf("tnc.Client.GetOutboundRequestByID: parse_data %v", err)
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("tnc.Client.GetOutboundRequestByID: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err = r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("tnc.Client.GetOutboundRequestByID: parse_response_err: %v", err)
		}
		return nil, fmt.Errorf("tnc.Client.GetOutboundRequestByID: failed code %s, message %s", errRes.Code, errRes.ErrorMessage)
	}
	if err = r.ParseResponseData(outboundRequest); err != nil {
		return nil, fmt.Errorf("tnc.Client.GetOutboundRequestByID: parse_response_data: %v", err)
	}
	return &outboundRequest, nil
}

// CancelOutboundRequest ...
func (c *Client) CancelOutboundRequest(requestID int) error {
	apiURL := c.getBaseURL() + fmt.Sprintf(apiPathCancelOutboundRequest, requestID)
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    apiURL,
			Method: http.MethodPost,
			Header: c.getRequestHeader(),
		},
	}
	msg, err := c.requestHttpViaNats(natsPayload)
	if err != nil {
		logger.Error("tnc.Client.CancelOutboundRequest - requestHttpViaNats", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return err
	}
	var (
		r      natsiomodel.NatsResponse
		errRes ErrRes
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return fmt.Errorf("tnc.Client.CancelOutboundRequest: parse_data %v", err)
	}
	res := r.Response
	if res == nil {
		return fmt.Errorf("tnc.Client.CancelOutboundRequest: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err = r.ParseResponseData(&errRes); err != nil {
			return fmt.Errorf("tnc.Client.CancelOutboundRequest: parse_response_err: %v", err)
		}
		return fmt.Errorf("tnc.Client.CancelOutboundRequest: failed code %s, message %s", errRes.Code, errRes.ErrorMessage)
	}
	return nil
}

func (c *Client) auth() (*authRes, error) {
	v := url.Values{}
	v.Add("realm", c.realm)
	v.Add("grant_type", "client_credentials")
	v.Add("client_id", c.clientID)
	v.Add("client_secret", c.clientSecret)

	body := v.Encode()
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	apiURL := baseURLAuthENVMapping[c.env] + fmt.Sprintf(apiPathAuth, c.realm)
	natsPayload := natsiomodel.NatsRequestHTTP{
		ResponseImmediately: true,
		Payload: natsiomodel.HTTPPayload{
			URL:    apiURL,
			Method: http.MethodPost,
			Data:   body,
			Header: header,
		},
	}
	msg, err := c.requestHttpViaNats(natsPayload)
	if err != nil {
		logger.Error("tnc.Client.auth - requestHttpViaNats", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return nil, err
	}
	var (
		r      natsiomodel.NatsResponse
		errRes ErrRes
		data   authRes
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return nil, fmt.Errorf("tnc.Client.auth: parse_data %v", err)
	}
	res := r.Response
	if res == nil {
		return nil, fmt.Errorf("tnc.Client.auth: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err = r.ParseResponseData(&errRes); err != nil {
			return nil, fmt.Errorf("tnc.Client.auth: parse_response_err: %v", err)
		}
		return nil, fmt.Errorf("tnc.Client.auth: failed code %s, message %s", errRes.Code, errRes.ErrorMessage)
	}
	if err = r.ParseResponseData(data); err != nil {
		return nil, fmt.Errorf("tnc.Client.auth: parse_response_data: %v", err)
	}
	return &data, nil
}

func (c *Client) getRequestHeader() map[string]string {
	m := map[string]string{
		"Content-Type": "application/json",
	}
	token, err := c.getToken()
	if err != nil {
		logger.Error("tnc.Client.getToken", logger.LogData{"err": err.Error()})
	} else {
		m["Authorization"] = fmt.Sprintf("Bearer %s", token)
	}
	return m
}

func (c *Client) requestHttpViaNats(data natsiomodel.NatsRequestHTTP) (*nats.Msg, error) {
	s := constant.NatsCommunicationSubjectRequestHTTP
	b := pjson.ToBytes(data)
	return c.natsClient.Request(s, b)
}

func (c *Client) getBaseURL() string {
	return baseURLENVMapping[c.env]
}

func (c *Client) getToken() (string, error) {
	if c.token != "" || c.tokenExpireAt.After(time.Now()) {
		return c.token, nil
	}
	data, err := c.auth()
	if err != nil {
		return "", err
	}
	c.token = data.AccessToken
	d := time.Duration(data.ExpiresIn) * time.Second
	if d.Minutes() > 30 {
		d -= 30 * time.Minute
	}
	c.tokenExpireAt = time.Now().Add(d)
	return c.token, nil
}
