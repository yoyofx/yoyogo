package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	defaultTransport *http.Transport
	BaseUrl          string
}

func NewClient() *Client {
	return &Client{defaultTransport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}}
}

// WithFormRequest 快速配置表单请求类型
func (c *Client) WithFormRequest(params map[string]interface{}) *Request {
	defaultHeader := make(http.Header)
	defaultHeader.Set("Content-Type", "application/x-www-form-urlencoded")
	request := &Request{header: defaultHeader, contentType: "application/x-www-form-urlencoded", params: params, timeout: 5}
	return request
}

func (c *Client) WithRequest() *Request {
	return &Request{header: make(http.Header), timeout: 5}
}

// WithFormRequest 快速配置表单请求类型
func (c *Client) WithJsonRequest(json string) *Request {
	defaultHeader := make(http.Header)
	defaultHeader.Set("Content-Type", "application/json")
	request := &Request{header: defaultHeader, contentType: "application/json", requestBody: []byte(json), timeout: 5}
	return request
}

// RunGet 执行Get请求
func (c *Client) Get(request *Request) (clientResp *Response, err error) {
	if request.errorRaw != "" {
		return nil, request.Error()
	}

	var defaultClient = &http.Client{}

	if request.cooJar != nil {
		defaultClient.Jar = request.cooJar
	}

	transport := c.defaultTransport
	transport.DisableKeepAlives = request.disableKeepAlives

	if request.skipHttps == true {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	defaultClient.Transport = http.RoundTripper(transport)

	req, err := http.NewRequest("GET", request.url, nil)
	if err != nil {
		return nil, err
	}
	timeSt := time.Now()
	// 超时设置
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(request.timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	if request.host != "" {
		req.Host = request.host
	}
	if request.header == nil {
		request.header = http.Header{}
	}
	request.header.Set("Cookie", "")
	req.Header = request.header

	resp, err := defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	for _, item := range resp.Cookies() {
		request.setCookieData(item.Name, item)
	}

	requestTime := time.Now().Sub(timeSt)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	clientResp = &Response{
		BodyRaw:     resp,
		Body:        body,
		RequestTime: requestTime,
		CookieData:  request.cookieData,
	}

	return clientResp, err
}

// RunGet 执行Post请求
func (c *Client) Post(request *Request) (clientResp *Response, err error) {
	clientResp = &Response{
		CookieData: request.cookieData,
	}

	if len(request.params) > 0 || len(request.files) > 0 {
		request.paraseParams()
	}

	if request.errorRaw != "" {
		return nil, request.Error()
	}

	var defaultClient = &http.Client{}
	if request.cooJar != nil {
		defaultClient.Jar = request.cooJar
	}

	transport := c.defaultTransport
	transport.DisableKeepAlives = request.disableKeepAlives
	if request.skipHttps == true {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	defaultClient.Transport = http.RoundTripper(transport)

	req, err := http.NewRequest("POST", request.url, bytes.NewBuffer(request.requestBody))
	if err != nil {
		return nil, err
	}

	timeSt := time.Now()
	// 超时设置
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(request.timeout)*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	request.header.Set("Cookie", "")
	req.Header = request.header
	req.Header.Set("Content-Type", request.contentType)

	if request.host != "" {
		req.Host = request.host
	}

	resp, err := defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	for _, item := range resp.Cookies() {
		request.setCookieData(item.Name, item)
	}

	requestTime := time.Now().Sub(timeSt)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))

	clientResp.BodyRaw = resp
	clientResp.Body = body
	clientResp.RequestTime = requestTime

	return clientResp, err
}

func (c *Client) Do(request *Request) (clientResp *Response, err error) {
	if request.method == "" {
		return nil, errors.New("this request is no method set.")
	}
	if !strings.HasPrefix(request.url, "http") {
		if c.BaseUrl == "" {
			return nil, errors.New("url don't have host and client don't config baseUrl please config")
		}
		request.url = c.BaseUrl + request.url
	}
	if request.method == "GET" {
		clientResp, err = c.Get(request)
	} else { // POST
		clientResp, err = c.Post(request)
	}
	return clientResp, err
}
