package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/config"
)

type Client struct {
	httpClient *http.Client
	config     *config.Config
	baseURL    string
}

func NewClient(config *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: config.HTTPTimeout,
		},
		config:  config,
		baseURL: config.BaseURL,
	}
}

// Get performs a GET request and decodes the response into the provided interface
func (c *Client) Get(endpoint string, v interface{}) error {
	fullURL, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return fmt.Errorf("joining URL paths: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	// Add API key to query parameters
	q := req.URL.Query()
	q.Add("api", c.config.APIKey)
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}
