#!/bin/bash

mkdir -p cmd/api-runner
mkdir -p internal/config
mkdir -p internal/api_clients/service1
mkdir -p internal/api_clients/service2
mkdir -p internal/models
mkdir -p internal/logging
mkdir -p pkg

touch cmd/api-runner/main.go
touch internal/config/config.go
touch internal/api_clients/service1/client.go
touch internal/api_clients/service2/client.go
touch internal/models/service1.go
touch internal/models/service2.go
touch internal/logging/logger.go
touch go.mod
touch .gitignore
touch README.md

echo "Project structure created."
