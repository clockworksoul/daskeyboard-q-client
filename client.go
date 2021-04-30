package qclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultBackendUrl = "http://localhost:27301"
	DefaultAPIPath    = "/api/1.0/"
	DefaultTimeout    = time.Minute
)

type Client struct {
	baseUrl *url.URL
	client  *http.Client
}

func (c *Client) APIURL(api string) *url.URL {
	return c.baseUrl.ResolveReference(&url.URL{Path: "signals"})
}

type ClientConfig struct {
	backendUrl    string
	apiPath       string
	clientTimeout time.Duration
}

type ClientOption func(*ClientConfig)

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &ClientConfig{
		backendUrl:    DefaultBackendUrl,
		apiPath:       DefaultAPIPath,
		clientTimeout: DefaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	u, err := url.Parse(c.backendUrl)
	if err != nil {
		return nil, err
	}
	u = u.ResolveReference(&url.URL{Path: c.apiPath})

	client := &Client{
		baseUrl: u,
		client:  &http.Client{Timeout: c.clientTimeout},
	}

	return client, nil
}

func WithTimeout(d time.Duration) ClientOption {
	return func(c *ClientConfig) {
		c.clientTimeout = d
	}
}

func (c *Client) Do(ctx context.Context, method, path string, input, output interface{}) error {
	u := c.baseUrl.ResolveReference(&url.URL{Path: path})
	var reqBody io.ReadCloser

	if input != nil {
		jreq, _ := json.Marshal(input)
		reqBody = io.NopCloser(bytes.NewReader(jreq))
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		msg, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%d %s", res.StatusCode, string(msg))
	}

	if output == nil {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(output); err != nil {
		return err
	}

	return err
}

func (c *Client) CreateSignal(ctx context.Context, s *SignalRequest) (*SignalResponse, error) {
	sres := &SignalResponse{}
	err := c.Do(ctx, "POST", "signals", s, sres)

	return sres, err
}

func (c *Client) DeleteSignalByID(ctx context.Context, id int) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("signals/%d", id), nil, nil)
}
