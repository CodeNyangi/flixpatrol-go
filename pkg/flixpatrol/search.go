package flixpatrol

import (
	"fmt"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// SearchService handles search-related operations for FlixPatrol API
type SearchService struct {
	client *api.Client
}

// NewSearchService creates a new SearchService instance
func NewSearchService(client *api.Client) *SearchService {
	return &SearchService{client: client}
}

// Search performs a search query on FlixPatrol API
func (s *SearchService) Search(query string, page int) ([]models.SearchResponse, error) {
	endpoint, err := s.buildSearchURL(query, page)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var searchResults []models.SearchResponse
	err = s.client.Get(endpoint, &searchResults)
	if err != nil {
		return nil, fmt.Errorf("performing search: %w", err)
	}

	return searchResults, nil
}

// buildSearchURL constructs the URL for the search endpoint
func (s *SearchService) buildSearchURL(query string, page int) (string, error) {
	endpoint, err := url.Parse("/search/")
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %w", err)
	}

	q := endpoint.Query()
	q.Add("query", query)
	q.Add("page", fmt.Sprintf("%d", page))
	endpoint.RawQuery = q.Encode()

	return endpoint.String(), nil
}
