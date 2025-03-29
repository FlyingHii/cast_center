package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Define a struct to represent the API response (adjust fields as needed)
type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Function to make the API call
func callAPI(apiEndpoint string) (*ApiResponse, error) {
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	// Decode the JSON response (example - adapt to your API's response format)
	var apiResponse ApiResponse
	// if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
	// 	return nil, fmt.Errorf("failed to decode response: %w", err)
	// }

	return &apiResponse, nil
}

func main() {
	// Load configuration
	config := LoadConfig()

	// Check if API endpoint is set
	if config.APIEndpoint == "" {
		log.Fatal("API_ENDPOINT environment variable not set")
	}

	// Create a ticker to run the API calls periodically
	ticker := time.NewTicker(5 * time.Second) // Adjust the interval as needed
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				// Call the API
				response, err := callAPI(config.APIEndpoint)
				if err != nil {
					log.Printf("API call failed: %v", err)
				} else {
					// Process the response (example)
					fmt.Printf("API Response: %+v\n", response)
				}
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	// Keep the program running (e.g., until a signal is received)
	// For a real application, you'd likely use a signal handler (e.g., to gracefully shut down)
	time.Sleep(30 * time.Second) // Run for 30 seconds
	done <- true
	fmt.Println("Runner stopped.")
}
