# Description

This project is a Go application that interacts with the Upwork API. It is designed to fetch and process data from Upwork, potentially for use in a larger "Cast Center" system.  This runner currently demonstrates fetching user information from the Upwork API.

# Prerequisites

*   **Go:** You need to have Go installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).
*   **Upwork API Credentials:** You need to create an Upwork API application and obtain your `consumer_key` and `consumer_secret`. Follow the instructions in the Upwork API documentation to create an application and retrieve these credentials.
*   **config.json:**  You need to have a `config.json` file in the `internal/api_clients/upwork/` directory (or the project root, depending on how you configure it) with your Upwork API credentials. An example `config.example.json` is provided.

# Setup

1.  **Clone the repository:**
    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Configure API Credentials:**
    *   Copy `internal/api_clients/upwork/config.example.json` to `internal/api_clients/upwork/config.json`.
    *   Edit `internal/api_clients/upwork/config.json` and replace the placeholder values with your actual `consumer_key` and `consumer_secret` from your Upwork API application.

# Usage

1.  **Run the `start.sh` script:**
    ```bash
    /bin/bash start.sh
    ```
    This script will:
    *   Build the Go application `cmd/api-runner/main.go` into an executable named `api-runner`.
    *   Run the `api-runner` executable.

# To-Do

*   **2025-03-29:** Initial setup of the project, including Go module initialization, basic structure, and Upwork API client integration.
*   **2025-03-29:** Updating the ID to get API key on upwork https://www.upwork.com/developer/keys/apply

# License

[MIT License](LICENSE) (or specify your desired license)
