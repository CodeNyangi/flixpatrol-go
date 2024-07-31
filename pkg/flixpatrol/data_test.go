package flixpatrol_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
	"github.com/popeyeGOEL/flixpatrol-go/pkg/flixpatrol"
	"github.com/stretchr/testify/assert"
)

// MockConfig is a test implementation of the Config interface
type MockConfig struct {
	apiKey      string
	baseURL     string
	httpTimeout time.Duration
}

func (m MockConfig) APIKey() string             { return m.apiKey }
func (m MockConfig) BaseURL() string            { return m.baseURL }
func (m MockConfig) HTTPTimeout() time.Duration { return m.httpTimeout }

func TestDataService_GetData(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Validate request method and path
			assert.Equal(t, "/data", r.URL.Path)
			assert.Equal(t, "GET", r.Method)

			// Validate query parameters
			query := r.URL.Query()
			assert.Equal(t, "4", query.Get("set"))
			assert.Equal(t, "656", query.Get("streaming"))
			assert.Equal(t, "4672", query.Get("region"))
			assert.Equal(t, "2020", query.Get("date"))
			assert.Equal(t, "1", query.Get("type"))
			assert.Equal(t, "test-api-key", query.Get("api"))

			// Mock response data
			mockResponse := &models.ApiResponse{
				List: []models.DataResponse{
					{
						Result:      1,
						ID:          12345,
						Name:        "Test Movie",
						URL:         "https://flixpatrol.com/title/test-movie",
						Premiere:    "2020-01-01", // 날짜는 문자열로 제공됨
						TypeID:      1,
						Type:        "Movie",
						CountryID:   1,
						Country:     "USA",
						CompanyID:   1,
						Company:     "Test Studio",
						Key:         "test_key",
						Note:        "Test note",
						Region:      1, // 지역 ID가 숫자일 경우
						Ranking:     1,
						RankingLast: 2,
						Value:       100,
						ValueLast:   90,
						ValueTotal:  1000,
						Countries:   50,
						Days:        30,
					},
				},
			}

			// Set content type and encode response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockResponse)
		}))
		defer server.Close()

		// Mock configuration
		mockConfig := MockConfig{
			apiKey:      "test-api-key",
			baseURL:     server.URL,
			httpTimeout: 10 * time.Second,
		}

		// Create client and data service
		client := api.NewClient(mockConfig)
		dataService := flixpatrol.NewDataService(client)

		// Set request parameters
		params := url.Values{
			"set":       {"4"},
			"streaming": {"656"},
			"region":    {"4672"},
			"date":      {"2020"},
			"type":      {"1"},
		}

		// Make the request
		apiResponse, err := dataService.GetData(params)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Validate the response data
		assert.Len(t, apiResponse.List, 1)
		assert.Equal(t, 1, apiResponse.List[0].Result)
		assert.Equal(t, 12345, apiResponse.List[0].ID)
		assert.Equal(t, "Test Movie", apiResponse.List[0].Name)
		assert.Equal(t, "https://flixpatrol.com/title/test-movie", apiResponse.List[0].URL)
		assert.Equal(t, "2020-01-01", apiResponse.List[0].Premiere) // 문자열 형식으로 확인
		assert.Equal(t, 1, apiResponse.List[0].TypeID)
		assert.Equal(t, "Movie", apiResponse.List[0].Type)
		assert.Equal(t, 1, apiResponse.List[0].CountryID)
		assert.Equal(t, "USA", apiResponse.List[0].Country)
		assert.Equal(t, 1, apiResponse.List[0].CompanyID)
		assert.Equal(t, "Test Studio", apiResponse.List[0].Company)
		assert.Equal(t, "test_key", apiResponse.List[0].Key)
		assert.Equal(t, "Test note", apiResponse.List[0].Note)
		assert.Equal(t, 1, apiResponse.List[0].Region) // 지역 ID가 숫자일 경우
		assert.Equal(t, 1, apiResponse.List[0].Ranking)
		assert.Equal(t, 2, apiResponse.List[0].RankingLast)
		assert.Equal(t, 100, apiResponse.List[0].Value)
		assert.Equal(t, 90, apiResponse.List[0].ValueLast)
		assert.Equal(t, 1000, apiResponse.List[0].ValueTotal)
		assert.Equal(t, 50, apiResponse.List[0].Countries)
		assert.Equal(t, 30, apiResponse.List[0].Days)
	})
}
