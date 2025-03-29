package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cast_center/internal/api_clients/upwork"

	"cast_center/internal/config" // Import the config package
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
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &apiResponse, nil
}

func main() {
	// Load configuration
	config := config.LoadConfig(upwork.CfgFile)

	// Check if API endpoint is set
	if config.APIEndpoint == "" {
		log.Fatal("API_ENDPOINT environment variable not set")
	}

	// Call the API
	upwork.GetUserInfo()

	fmt.Println("Runner stopped.")
}
