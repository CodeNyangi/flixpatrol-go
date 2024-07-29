package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/config"
	"github.com/popeyeGOEL/flixpatrol-go/pkg/flixpatrol"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	// Create a new client
	client := api.NewClient(cfg)

	// Initialize services
	dataService := flixpatrol.NewDataService(client)
	demographicsService := flixpatrol.NewDemographicsService(client)
	preferencesService := flixpatrol.NewPreferencesService(client)
	searchService := flixpatrol.NewSearchService(client)
	titleService := flixpatrol.NewTitleService(client)
	trendingService := flixpatrol.NewTrendingService(client)

	// Example: Get data
	dataParams := url.Values{
		"set":       {"4"},
		"streaming": {"656"},
		"region":    {"4672"},
		"date":      {"2020"},
		"type":      {"1"},
	}
	data, err := dataService.GetData(dataParams)
	if err != nil {
		return fmt.Errorf("getting data: %w", err)
	}
	fmt.Printf("Data: %+v\n", data)

	// Example: Get demographics
	demographicsParams := url.Values{
		"region": {"4672"},
		"date":   {"2020"},
	}
	demographics, err := demographicsService.GetDemographics(demographicsParams)
	if err != nil {
		return fmt.Errorf("getting demographics: %w", err)
	}
	fmt.Printf("Demographics: %+v\n", demographics)

	// Example: Get preferences
	preferencesParams := url.Values{
		"set":       {"33"},
		"streaming": {"656"},
		"region":    {"4672"},
		"date":      {"2020-10"},
	}
	preferences, err := preferencesService.GetPreferences(preferencesParams)
	if err != nil {
		return fmt.Errorf("getting preferences: %w", err)
	}
	fmt.Printf("Preferences: %+v\n", preferences)

	// Example: Search
	searchResults, err := searchService.Search("star wars", 1)
	if err != nil {
		return fmt.Errorf("searching: %w", err)
	}
	fmt.Printf("Search Results: %+v\n", searchResults)

	// Example: Get title
	titleOpts := &flixpatrol.TitleOptions{
		Set:       1,
		Streaming: 656,
		Region:    0,
		Date:      "2020-07-24",
	}
	title, err := titleService.GetTitle(89624, titleOpts)
	if err != nil {
		return fmt.Errorf("getting title: %w", err)
	}
	fmt.Printf("Title: %+v\n", title)

	// Example: Get trending
	trendingOpts := &flixpatrol.TrendingOptions{
		Region: 4672,
		Date:   "2021-03",
	}
	trending, err := trendingService.GetTrending(trendingOpts)
	if err != nil {
		return fmt.Errorf("getting trending: %w", err)
	}
	fmt.Printf("Trending: %+v\n", trending)

	return nil
}
