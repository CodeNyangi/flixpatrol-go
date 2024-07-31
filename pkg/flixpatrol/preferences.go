package flixpatrol

import (
	"fmt"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
)

// PreferencesService handles preferences-related operations for FlixPatrol API
type PreferencesService struct {
	client *api.Client
}

// NewPreferencesService creates a new PreferencesService instance
func NewPreferencesService(client *api.Client) *PreferencesService {
	return &PreferencesService{client: client}
}

// GetPreferences retrieves preferences data from FlixPatrol API based on provided parameters
func (s *PreferencesService) GetPreferences(params url.Values) (*models.PreferencesResponse, error) {
	endpoint, err := s.buildPreferencesURL(params)
	if err != nil {
		return nil, fmt.Errorf("building URL: %w", err)
	}

	var preferences models.PreferencesResponse
	err = s.client.Get(endpoint, &preferences)
	if err != nil {
		return nil, fmt.Errorf("getting preferences data: %w", err)
	}

	return &preferences, nil
}

// buildPreferencesURL constructs the URL for the preferences endpoint
func (s *PreferencesService) buildPreferencesURL(params url.Values) (string, error) {
	endpoint, err := url.Parse("/preferences/")
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
