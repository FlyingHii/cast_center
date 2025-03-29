package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPRequest performs an HTTP request and returns the response body as a byte slice.
func HTTPRequest(method, url string, body interface{}, headers map[string]string) ([]byte, int, error) {
	client := &http.Client{}
	var req *http.Request
	var err error

	// Prepare the request body
	var bodyReader *bytes.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	// Create the request
	if bodyReader != nil {
		req, err = http.NewRequest(method, url, bodyReader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading response body: %w", err)
	}

	return respBody, resp.StatusCode, nil
}

// Get performs an HTTP GET request.
func Get(url string, headers map[string]string) ([]byte, int, error) {
	return HTTPRequest(http.MethodGet, url, nil, headers)
}

// Post performs an HTTP POST request.
func Post(url string, body interface{}, headers map[string]string) ([]byte, int, error) {
	return HTTPRequest(http.MethodPost, url, body, headers)
}

// Put performs an HTTP PUT request.
func Put(url string, body interface{}, headers map[string]string) ([]byte, int, error) {
	return HTTPRequest(http.MethodPut, url, body, headers)
}
