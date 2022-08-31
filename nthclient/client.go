package nthclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-multierror"
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

func (c *Client) prepareRequest(ctx context.Context, urlString string) (*http.Request, error) {
	if urlString == "" {
		urlObject := url.URL{
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
		urlString = urlObject.String()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlString, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to construct request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.settings.UserAgent)
	req.Header.Set("Accept-Language", c.settings.Language)
	return req, nil
}

func (c *Client) getEncryptedBody(ctx context.Context) ([]byte, error) {
	var resErr error
	targets := append([]string{"", "", ""}, c.settings.BackupDomains...)
	for _, urlString := range targets {
		ctx1, cl := context.WithTimeout(ctx, c.settings.Timeout)
		defer cl()
		req, err := c.prepareRequest(ctx1, urlString)
		if err != nil {
			resErr = multierror.Append(resErr, fmt.Errorf("unable to prepare request: %w", err))
			continue
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			resErr = multierror.Append(resErr, fmt.Errorf("API request failed: %w", err))
			continue
		}
		defer cleanupBody(resp.Body)

		rd := &io.LimitedReader{
			R: resp.Body,
			N: readLimit,
		}

		if resp.StatusCode != http.StatusOK {
			respBytes, _ := io.ReadAll(rd)
			resErr = multierror.Append(resErr,
				fmt.Errorf("bad status code from API: code = %d, body = %q", resp.StatusCode, string(respBytes)),
			)
			continue
		}

		respBytes, err := io.ReadAll(rd)
		if err != nil {
			resErr = multierror.Append(resErr, fmt.Errorf("API response read failed: %w", err))
			continue
		}

		payload, err := VerifyResponse(string(respBytes), c.settings.PublicKey)
		if err != nil {
			resErr = multierror.Append(resErr, fmt.Errorf("API response verification failed: %w", err))
			continue
		}

		decrypted, err := Decrypt(string(payload), c.settings.JSONSeed)
		if err != nil {
			resErr = multierror.Append(resErr, fmt.Errorf("payload decryption failed: %w", err))
			continue
		}

		return decrypted, nil
	}
	return nil, resErr
}

func (c *Client) GetServerConfig(ctx context.Context) ([]byte, error) {
	return c.getEncryptedBody(ctx)
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
