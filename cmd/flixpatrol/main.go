package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/popeyeGOEL/flixpatrol-go/internal/api"
	"github.com/popeyeGOEL/flixpatrol-go/internal/config"
	"github.com/popeyeGOEL/flixpatrol-go/internal/models"
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
	preferencesService := flixpatrol.NewPreferencesService(client)

	// Example: Get data
	dataParams := url.Values{
		"set":       {"4"},
		"streaming": {"656"},
		"region":    {"4672"},
		"date":      {"2024"},
		"type":      {"1"},
	}
	dataResponse, err := dataService.GetData(dataParams)
	if err != nil {
		return fmt.Errorf("getting data: %w", err)
	}
	fmt.Printf("Data: %+v\n", dataResponse)

	// Write data to CSV
	if err := writeDataToCSV("data.csv", dataResponse.List); err != nil {
		return fmt.Errorf("writing data to CSV: %w", err)
	}

	// Example: Get preferences
	preferencesParams := url.Values{
		"set":       {"1"},
		"streaming": {"656"},
		"region":    {"4672"},
		"date":      {"2024"},
	}
	preferences, err := preferencesService.GetPreferences(preferencesParams)
	if err != nil {
		return fmt.Errorf("getting preferences: %w", err)
	}
	fmt.Printf("Preferences: %+v\n", preferences)

	// Write preferences to CSV
	if err := writePreferencesToCSV("preferences.csv", preferences); err != nil {
		return fmt.Errorf("writing preferences to CSV: %w", err)
	}

	return nil
}

func writeDataToCSV(filename string, data []models.DataResponse) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Result", "ID", "Name", "URL", "Premiere", "TypeID", "Type", "CountryID", "Country", "CompanyID", "Company", "Key", "Note", "Region", "Ranking", "RankingLast", "Value", "ValueLast", "ValueTotal", "Countries", "Days"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("writing header to CSV: %w", err)
	}

	// Write data
	for _, item := range data {
		record := []string{
			fmt.Sprint(item.Result),
			fmt.Sprint(item.ID),
			item.Name,
			item.URL,
			item.Premiere,
			fmt.Sprint(item.TypeID),
			item.Type,
			fmt.Sprint(item.CountryID),
			item.Country,
			fmt.Sprint(item.CompanyID),
			item.Company,
			item.Key,
			item.Note,
			fmt.Sprint(item.Region),
			fmt.Sprint(item.Ranking),
			fmt.Sprint(item.RankingLast),
			fmt.Sprint(item.Value),
			fmt.Sprint(item.ValueLast),
			fmt.Sprint(item.ValueTotal),
			fmt.Sprint(item.Countries),
			fmt.Sprint(item.Days),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("writing record to CSV: %w", err)
		}
	}

	return nil
}

func writePreferencesToCSV(filename string, preferences *models.PreferencesResponse) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating CSV file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Result", "ID", "Name", "GID", "Group", "Value", "Share"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("writing header to CSV: %w", err)
	}

	// Write data
	for _, item := range preferences.List {
		record := []string{
			fmt.Sprint(item.Result),
			fmt.Sprint(item.ID),
			item.Name,
			fmt.Sprint(item.GID),
			item.Group,
			fmt.Sprint(item.Value),
			fmt.Sprintf("%.2f", item.Share),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("writing record to CSV: %w", err)
		}
	}

	return nil
}
