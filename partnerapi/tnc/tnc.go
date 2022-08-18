package tnc

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/natsio"
	"github.com/Selly-Modules/natsio/model"
	"github.com/nats-io/nats.go"

	"github.com/Selly-Modules/3pl/constant"
	"github.com/Selly-Modules/3pl/util/pjson"
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
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    apiURL,
			Method: http.MethodPost,
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
		r       model.CommunicationHttpResponse
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
	item := &dataRes[0]
	e := item.Error
	if e != nil {
		return nil, fmt.Errorf("tnc.Client.CreateOutboundRequest: failed, code %s - message %s", e.Code, e.ErrorMessage)
	}

	return item, err
}

// UpdateOutboundRequestLogisticInfo ...
func (c *Client) UpdateOutboundRequestLogisticInfo(p UpdateORLogisticInfoPayload) error {
	apiURL := c.getBaseURL() + fmt.Sprintf(apiPathUpdateLogisticInfoOutboundRequest, p.OrID)
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    apiURL,
			Method: http.MethodPost,
			Header: c.getRequestHeader(),
		},
	}
	msg, err := c.requestHttpViaNats(natsPayload)
	if err != nil {
		logger.Error("tnc.Client.UpdateOutboundRequestLogisticInfo - requestHttpViaNats", logger.LogData{
			"err":     err.Error(),
			"payload": natsPayload,
		})
		return err
	}
	var (
		r      model.CommunicationHttpResponse
		errRes ErrRes
	)
	if err = pjson.Unmarshal(msg.Data, &r); err != nil {
		return fmt.Errorf("tnc.Client.UpdateOutboundRequestLogisticInfo: parse_data %v", err)
	}
	res := r.Response
	if res == nil {
		return fmt.Errorf("tnc.Client.UpdateOutboundRequestLogisticInfo: empty_response")
	}
	if res.StatusCode >= http.StatusBadRequest {
		if err = r.ParseResponseData(&errRes); err != nil {
			return fmt.Errorf("tnc.Client.UpdateOutboundRequestLogisticInfo: parse_response_err: %v", err)
		}
		return fmt.Errorf("tnc.Client.UpdateOutboundRequestLogisticInfo: failed code %s, message %s", errRes.Code, errRes.ErrorMessage)
	}
	return nil
}

// GetOutboundRequestByID ...
func (c *Client) GetOutboundRequestByID(requestID int) (*OutboundRequestInfo, error) {
	apiURL := c.getBaseURL() + fmt.Sprintf(apiPathGetOutboundRequest, requestID)
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
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
		r               model.CommunicationHttpResponse
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
func (c *Client) CancelOutboundRequest(requestID int, note string) error {
	apiURL := c.getBaseURL() + fmt.Sprintf(apiPathCancelOutboundRequest, requestID)
	data := map[string]string{"note": note}
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
			URL:    apiURL,
			Method: http.MethodPost,
			Header: c.getRequestHeader(),
			Data:   pjson.ToJSONString(data),
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
		r      model.CommunicationHttpResponse
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
	natsPayload := model.CommunicationRequestHttp{
		ResponseImmediately: true,
		Payload: model.HttpRequest{
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
		r      model.CommunicationHttpResponse
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

func (c *Client) requestHttpViaNats(data model.CommunicationRequestHttp) (*nats.Msg, error) {
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
