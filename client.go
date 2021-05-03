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
	apiKey  string
	baseUrl *url.URL
	client  *http.Client
}

type ClientConfig struct {
	backendUrl    string
	apiPath       string
	apiKey        string
	clientTimeout time.Duration
}

type ClientOption func(*ClientConfig)

// NewClient is used to retrieve a new Q client. It accepts zero or more
// options returned by the various With* functions.
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

// WithAPIPath can be used to set the API key for interacting with Q Cloud. When set the
// API key will be included in a X-API-KEY request header in all requests.
// Default: None.
func WithAPIKey(s string) ClientOption {
	return func(c *ClientConfig) {
		c.apiKey = s
	}
}

// WithAPIPath can be used to set the API path. You shouldn't need to ever change this.
// Default: /api/1.0/.
func WithAPIPath(s string) ClientOption {
	return func(c *ClientConfig) {
		c.apiPath = s
	}
}

// WithBackendURL can be used to set the base URL for the target API.
// Default: http://localhost:27301.
func WithBackendURL(s string) ClientOption {
	return func(c *ClientConfig) {
		c.backendUrl = s
	}
}

// WithTimeout can be used to set the timeout for the underlying HTTP client.
// Default: 1 minute.
func WithTimeout(d time.Duration) ClientOption {
	return func(c *ClientConfig) {
		c.clientTimeout = d
	}
}

// Do is a low-level method that can be used to execute a raw request to the
// service API. Generally you wouldn't use this directly, instead using the
// various Create|Delete|Get methods. The "method" parameter is the HTTP method
// to use (GET, POST, etc); "path" is the API path ("signals"); "input" is an
// optional struct representing the request body; "output" is a struct that an
// optional response body will be unmarshalled into.
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

	if c.apiKey != "" {
		req.Header.Add("X-API-KEY", c.apiKey)
	}

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

// CreateSignal can be used to generate a message to a Q-enabled device. It may
// contain lighting color and effect information as well as a message for a
// human.
func (c *Client) CreateSignal(ctx context.Context, s *SignalRequest) (*SignalResponse, error) {
	sres := &SignalResponse{}
	err := c.Do(ctx, "POST", "signals", s, sres)

	return sres, err
}

// DeleteSignalByID can be used to delete a signal using the signal id.
func (c *Client) DeleteSignalByID(ctx context.Context, id int) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("signals/%d", id), nil, nil)
}

// DeleteSignalByZoneID can be used to retrieve signals by a zone ID. See
// https://www.daskeyboard.io/q-zone-id-explanation/ for more information.
func (c *Client) DeleteSignalByZoneID(ctx context.Context, productId string, zoneID ZoneID) error {
	return c.Do(ctx, "DELETE", fmt.Sprintf("signals/pid/%s/zoneId/%s", productId, zoneID), nil, nil)
}

// GetShadowsByProductID lists the shadows, the list of the most recent
// signals for each zone.
func (c *Client) GetShadowsByProductID(ctx context.Context, productId string) ([]*SignalResponse, error) {
	srs := []*SignalResponse{}

	return srs, c.Do(ctx, "GET", fmt.Sprintf("signals/pid/%s/shadows", productId), nil, srs)
}

// GetShadowsByZoneID
func (c *Client) GetShadowsByZoneID(ctx context.Context, productId string, zoneID ZoneID) ([]*SignalResponse, error) {
	srs := []*SignalResponse{}

	return srs, c.Do(ctx, "GET", fmt.Sprintf("signals/pid/%s/zoneId/%s", productId, zoneID), nil, srs)
}

// GetShadows gets all available shadows.
func (c *Client) GetShadows(ctx context.Context) ([]*SignalResponse, error) {
	srs := []*SignalResponse{}
	return srs, c.Do(ctx, "GET", "signals/shadows", nil, &srs)
}

// GetSignals fetches a list of signals using pagination. It is only supported
//in Q Cloud.
func (c *Client) GetSignals(ctx context.Context, page, signalsPerPage int, sortBy string, ascending bool) (*SignalResponsePage, error) {
	var srp = &SignalResponsePage{}
	var asc string

	if ascending {
		asc = "ASC"
	} else {
		asc = "DESC"
	}

	path := fmt.Sprintf("signals?page=%d&size=%d&sort=%s,%s", page, signalsPerPage, sortBy, asc)

	return srp, c.Do(ctx, "GET", path, nil, &page)
}
