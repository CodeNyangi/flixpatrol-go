package flixpatrol

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// TrendingService handles trending-related operations for FlixPatrol API
type TrendingService struct {
	client *api.Client
}

// NewTrendingService creates a new TrendingService instance
func NewTrendingService(client *api.Client) *TrendingService {
	return &TrendingService{client: client}
}

// GetTrending retrieves trending data
func (s *TrendingService) GetTrending(opts *TrendingOptions) ([]models.TrendingResponse, error) {
	endpoint, err := s.buildTrendingURL(opts)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var trendingResponse []models.TrendingResponse
	err = s.client.Get(endpoint, &trendingResponse)
	if err != nil {
		return nil, fmt.Errorf("fetching trending data: %w", err)
	}

	return trendingResponse, nil
}

// TrendingOptions represents optional parameters for the trending endpoint
type TrendingOptions struct {
	Region int
	Date   string
	Page   int
}

// buildTrendingURL constructs the URL for the trending endpoint
func (s *TrendingService) buildTrendingURL(opts *TrendingOptions) (string, error) {
	endpoint, err := url.Parse("/trending/")
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %w", err)
	}

	q := endpoint.Query()

	if opts != nil {
		if opts.Region != 0 {
			q.Add("region", strconv.Itoa(opts.Region))
		}
		if opts.Date != "" {
			q.Add("date", opts.Date)
		}
		if opts.Page != 0 {
			q.Add("page", strconv.Itoa(opts.Page))
		}
	}

	endpoint.RawQuery = q.Encode()
	return endpoint.String(), nil
}
