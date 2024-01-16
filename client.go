package unisat_sdk_go

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeAndTwo/unisat-sdk-go/constant"
	"github.com/ThreeAndTwo/unisat-sdk-go/types"
	"github.com/imroc/req/v3"
	"strings"
)

type Client struct {
	httpClient *req.Client
	baseURL    string
	authToken  string
}

func NewClient(url, authToken string) *Client {
	if url == "" {
		url = constant.DefaultBaseURL
	}

	httpClient := req.C()
	httpClient.SetCommonHeader("Authorization", "Bearer "+authToken)
	return &Client{
		httpClient: httpClient,
		baseURL:    url,
	}
}

func (c *Client) GetUrl() string {
	return c.baseURL
}

func (c *Client) SetUrl(url string) {
	c.baseURL = url
}

func (c *Client) GetAuthToken() string {
	return c.authToken
}

func (c *Client) SetAuthToken(authToken string) {
	c.authToken = authToken
}

func (c *Client) GetBRC20Status(start, limit int) (*types.Brc20StatusRes, error) {
	url := c.baseURL + "/brc20/status"
	if limit == 0 {
		limit = constant.DefaultLimit
	}

	resp, err := c.httpClient.R().
		SetQueryParam("start", fmt.Sprint(start)).
		SetQueryParam("limit", fmt.Sprint(limit)).
		Get(url)
	if err != nil {
		return nil, err
	}

	v, err := c.dealResponse(new(types.Brc20StatusRes), resp, err)
	if err != nil {
		return nil, err
	}
	return v.(*types.Brc20StatusRes), nil
}

func (c *Client) GetBrc20TickerInfo(ticker string) (*types.Brc20TickerInfoRes, error) {
	url := c.baseURL + "/brc20/" + ticker + "/info"
	resp, err := c.httpClient.R().Get(url)
	if err != nil {
		return nil, err
	}

	v, err := c.dealResponse(new(types.Brc20TickerInfoRes), resp, err)
	if err != nil {
		return nil, err
	}
	return v.(*types.Brc20TickerInfoRes), nil
}

func (c *Client) GetBrc20Holders(ticker string, start, limit int) (*types.Brc20HoldersRes, error) {
	url := c.baseURL + "/brc20/" + ticker + "/holders"
	if limit == 0 {
		limit = constant.DefaultLimit
	}

	resp, err := c.httpClient.R().
		SetQueryParam("start", fmt.Sprint(start)).
		SetQueryParam("limit", fmt.Sprint(limit)).
		Get(url)
	if err != nil {
		return nil, err
	}

	v, err := c.dealResponse(new(types.Brc20HoldersRes), resp, err)
	if err != nil {
		return nil, err
	}
	return v.(*types.Brc20HoldersRes), nil
}

func (c *Client) GetAddressSummary(address string, start, limit int) (*types.Brc20SummaryRes, error) {
	url := c.baseURL + "/address/" + address + "/brc20/summary"
	if limit == 0 {
		limit = constant.DefaultLimit
	}
	resp, err := c.httpClient.R().
		SetQueryParam("start", fmt.Sprint(start)).
		SetQueryParam("limit", fmt.Sprint(limit)).
		Get(url)
	if err != nil {
		return nil, err
	}

	v, err := c.dealResponse(new(types.Brc20SummaryRes), resp, err)
	if err != nil {
		return nil, err
	}
	return v.(*types.Brc20SummaryRes), nil
}

func (c *Client) dealResponse(data interface{}, r *req.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	if r.Response.StatusCode != 200 {
		return nil, fmt.Errorf("err StatusCode %d, body size %d", r.Response.StatusCode, r.Response.ContentLength)
	}
	body := r.String()
	if strings.HasPrefix(body, "[") {
		body = body[1 : len(body)-1]
	}
	err = json.Unmarshal([]byte(body), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
