# FlixPatrol Go Client

This is an unofficial Go client for the FlixPatrol API. It provides easy access to various FlixPatrol services including data retrieval, demographics, preferences, search, title information, and trending content.

## Installation

To install the FlixPatrol Go client, use `go get`:

```bash
go get github.com/popeyeGOEL/flixpatrol-go
```

## Usage

Here's a quick example of how to use the FlixPatrol Go client:

```go
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
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create a new client
	client := api.NewClient(cfg)

	// Initialize a service, e.g., search service
	searchService := flixpatrol.NewSearchService(client)

	// Use the service
	searchResults, err := searchService.Search("star wars", 1)
	if err != nil {
		log.Fatalf("Error searching: %v", err)
	}

	fmt.Printf("Search Results: %+v\n", searchResults)
}
```

## Available Services

- Data Service
- Demographics Service
- Preferences Service
- Search Service
- Title Service
- Trending Service

Each service provides methods to interact with different endpoints of the FlixPatrol API.

## Configuration

The client requires configuration to be set up. Create a `config.yaml` file in the project root with the following structure:

```yaml
api:
  base_url: "https://api.flixpatrol.com"
  key: "your-api-key-here"
```

## Testing

To run tests:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This is an unofficial client and is not affiliated with, maintained, authorized, endorsed, or sponsored by FlixPatrol.
