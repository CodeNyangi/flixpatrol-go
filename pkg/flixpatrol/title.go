package flixpatrol

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// TitleService handles title-related operations for FlixPatrol API
type TitleService struct {
	client *api.Client
}

// NewTitleService creates a new TitleService instance
func NewTitleService(client *api.Client) *TitleService {
	return &TitleService{client: client}
}

// GetTitle retrieves information about a specific title
func (s *TitleService) GetTitle(id int, opts *TitleOptions) (*models.TitleResponse, error) {
	endpoint, err := s.buildTitleURL(id, opts)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var titleResponse models.TitleResponse
	err = s.client.Get(endpoint, &titleResponse)
	if err != nil {
		return nil, fmt.Errorf("fetching title data: %w", err)
	}

	return &titleResponse, nil
}

// TitleOptions represents optional parameters for the title endpoint
type TitleOptions struct {
	Set       int
	Streaming int
	Region    int
	Date      string
	Top25     bool
	Page      int
}

// buildTitleURL constructs the URL for the title endpoint
func (s *TitleService) buildTitleURL(id int, opts *TitleOptions) (string, error) {
	endpoint, err := url.Parse("/title/")
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %w", err)
	}

	q := endpoint.Query()
	q.Add("id", strconv.Itoa(id))

	if opts != nil {
		if opts.Set != 0 {
			q.Add("set", strconv.Itoa(opts.Set))
		}
		if opts.Streaming != 0 {
			q.Add("streaming", strconv.Itoa(opts.Streaming))
		}
		if opts.Region != 0 {
			q.Add("region", strconv.Itoa(opts.Region))
		}
		if opts.Date != "" {
			q.Add("date", opts.Date)
		}
		if opts.Top25 {
			q.Add("top25", "1")
		}
		if opts.Page != 0 {
			q.Add("page", strconv.Itoa(opts.Page))
		}
	}

	endpoint.RawQuery = q.Encode()
	return endpoint.String(), nil
}
