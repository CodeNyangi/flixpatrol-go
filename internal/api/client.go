package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/popeyeGOEL/flixpatrol-go/internal/config"
)

type Client struct {
	httpClient *http.Client
	config     config.Config
	baseURL    string
}

func NewClient(config config.Config) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second, // 타임아웃을 30초로 증가
		},
		config:  config,
		baseURL: config.BaseURL(),
	}
}

// Get performs a GET request and decodes the response into the provided interface
func (c *Client) Get(fullURL string, v interface{}) error {
	// concat base URL and endpoint URL
	u, err := url.Parse(c.baseURL + fullURL)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}

	// 끝의 슬래시 제거
	u.Path = u.Path[:len(u.Path)-1]

	// Add API key to query parameters
	q := u.Query()
	q.Add("api", c.config.APIKey())
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

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
