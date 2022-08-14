package nthclient

import "net/http"

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
