package nthclient

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	log.Printf("url = %q", url.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to construct request: %w", err)
	}
	return req, nil
}

func (c *Client) GetServerConfig(ctx context.Context) ([]byte, error) {
	_, err := c.prepareRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare request: %w", err)
	}
	return nil, nil
}
