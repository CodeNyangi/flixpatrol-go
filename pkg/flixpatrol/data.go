package flixpatrol

import (
	"fmt"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// DataService handles data-related operations for FlixPatrol API
type DataService struct {
	client *api.Client
}

// NewDataService creates a new DataService instance
func NewDataService(client *api.Client) *DataService {
	return &DataService{client: client}
}

func (s *DataService) GetData(params url.Values) (*models.ApiResponse, error) {
	endpoint, err := s.buildDataURL(params)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var response models.ApiResponse
	err = s.client.Get(endpoint, &response)
	if err != nil {
		return nil, fmt.Errorf("getting data: %w", err)
	}

	return &response, nil
}

// buildDataURL constructs the URL for the data endpoint
func (s *DataService) buildDataURL(params url.Values) (string, error) {
	endpoint, err := url.Parse("/data/")
	if err != nil {
		return "", fmt.Errorf("parsing base URL: %w", err)
	}

	query := endpoint.Query()
	for key, values := range params {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	endpoint.RawQuery = query.Encode()

	return endpoint.String(), nil
}
