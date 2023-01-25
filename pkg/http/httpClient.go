package http

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var clientMap sync.Map

type HttpClient struct {
	httpClient *http.Client
}

type HttpRequest struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Body    string
}

type HttpResponse struct {
	Body string
}

func newHttpClient(timeout int64) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives:   false,
		MaxIdleConns:        50000,
		MaxIdleConnsPerHost: 50000,
		IdleConnTimeout:     60 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}
	return &http.Client{Transport: transport, Timeout: time.Duration(timeout) * time.Millisecond}
}

func (client *HttpClient) Post(httpRequest *HttpRequest) (*HttpResponse, error) {
	u, err := url.Parse(httpRequest.Url)
	if err != nil {
		return nil, errors.Wrapf(err, "parse rawURL `%s` err", httpRequest.Url)
	}
	if len(httpRequest.Params) > 0 {
		query := u.Query()
		for k, v := range httpRequest.Params {
			query.Add(k, v)
		}
		u.RawQuery = query.Encode()
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(httpRequest.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if len(httpRequest.Headers) > 0 {
		for k, v := range httpRequest.Headers {
			req.Header.Set(k, v)
		}
	}
	httpResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "httpClient rawURL `%s` err", httpRequest.Url)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("do [%s %s] return code: %d message: %s", httpResp.Request.Method, httpRequest.Url, httpResp.StatusCode, "")
	}
	bs, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.Errorf("do [%s %s] ReadAll: %d message: %s", httpResp.Request.Method, httpRequest.Url, httpResp.StatusCode, err.Error())
	}

	return &HttpResponse{Body: string(bs)}, nil
}

func GetHttpClient(timeout int64) *HttpClient {
	httpClient, ok := clientMap.Load(timeout)

	//并发高的情况下，可能导致httpClient初始化两次
	if !ok || httpClient == nil {
		httpClient := newHttpClient(timeout)
		newHttpClient := &HttpClient{httpClient: httpClient}
		clientMap.Store(timeout, newHttpClient)
		return newHttpClient
	}
	return httpClient.(*HttpClient)
}
