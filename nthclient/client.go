package nthclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client for nthLink API backend, capable of retrieving server configuration from it.
type Client struct {
	settings   *Settings
	httpClient *http.Client
}

// Creates new client
func New() *Client {
	return &Client{
		settings:   DefaultSettings,
		httpClient: &http.Client{},
	}
}

// Copies Client
func (c *Client) Clone() *Client {
	newClient := *c
	return &newClient
}

// Creates new Client with specified http.Client
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	newClient := c.Clone()
	newClient.httpClient = httpClient
	return newClient
}

// Creates new Client with specified Settings
func (c *Client) WithSettings(settings *Settings) *Client {
	newClient := c.Clone()
	newClient.settings = settings
	return newClient
}

func (c *Client) prepareRequest(ctx context.Context) (*http.Request, error) {
	url := url.URL{
		Scheme: "https",
		Host:   CalculateAPIHostname(c.settings.DomainSeed, c.settings.TLD),
		Path:   ConfigRoutePath,
		RawQuery: url.Values{
			"key":        []string{c.settings.PlatformKey},
			"id":         []string{c.settings.ID},
			"lang":       []string{c.settings.Language},
			"appVersion": []string{c.settings.AppVersion},
		}.Encode(),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to construct request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.settings.UserAgent)
	req.Header.Set("Accept-Language", c.settings.Language)
	return req, nil
}

func (c *Client) GetServerConfig(ctx context.Context) ([]byte, error) {
	req, err := c.prepareRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer cleanupBody(resp.Body)

	rd := &io.LimitedReader{
		R: resp.Body,
		N: readLimit,
	}

	respBytes, err := io.ReadAll(rd)
	if err != nil {
		return nil, fmt.Errorf("API response read failed: %w", err)
	}

	parts := strings.SplitN(string(respBytes), "*-*", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("data was not found in the response. parts found: %d", len(parts))
	}
	return respBytes, nil
}

const readLimit int64 = 128 * 1024

// Does cleanup of HTTP response in order to make it reusable by keep-alive
// logic of HTTP client
func cleanupBody(body io.ReadCloser) {
	io.Copy(io.Discard, &io.LimitedReader{
		R: body,
		N: readLimit,
	})
	body.Close()
}
