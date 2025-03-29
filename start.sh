#!/bin/bash

# Build the Go program
go build -o api-runner cmd/api-runner/main.go

# Run the program, setting the API_ENDPOINT environment variable
API_ENDPOINT="https://api.example.com/data" # Replace with your API endpoint
./api-runner
