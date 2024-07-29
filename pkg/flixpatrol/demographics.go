package flixpatrol

import (
	"fmt"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// DemographicsService handles demographics-related operations for FlixPatrol API
type DemographicsService struct {
	client *api.Client
}

// NewDemographicsService creates a new DemographicsService instance
func NewDemographicsService(client *api.Client) *DemographicsService {
	return &DemographicsService{client: client}
}

// GetDemographics retrieves demographics data from FlixPatrol API based on provided parameters
func (s *DemographicsService) GetDemographics(params url.Values) ([]models.DemographicsResponse, error) {
	endpoint, err := s.buildDemographicsURL(params)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var demographics []models.DemographicsResponse
	err = s.client.Get(endpoint, &demographics)
	if err != nil {
		return nil, fmt.Errorf("getting demographics data: %w", err)
	}

	return demographics, nil
}

// buildDemographicsURL constructs the URL for the demographics endpoint
func (s *DemographicsService) buildDemographicsURL(params url.Values) (string, error) {
	endpoint, err := url.Parse("/demographics/")
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
