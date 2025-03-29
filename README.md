# Cast Center API Runner

## Description

This project is a Go application that interacts with the Upwork API. It is designed to fetch and process data from Upwork, potentially for use in a larger "Cast Center" system.  This runner currently demonstrates fetching user information from the Upwork API.

## Prerequisites

*   **Go:** You need to have Go installed on your system. You can download it from [https://go.dev/dl/](https://go.dev/dl/).
*   **Upwork API Credentials:** You need to create an Upwork API application and obtain your `consumer_key` and `consumer_secret`. Follow the instructions in the Upwork API documentation to create an application and retrieve these credentials.
*   **config.json:**  You need to have a `config.json` file in the `internal/api_clients/upwork/` directory (or the project root, depending on how you configure it) with your Upwork API credentials. An example `config.example.json` is provided.

## Setup

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

## Usage

1.  **Run the `start.sh` script:**
    ```bash
    /bin/bash start.sh
    ```
    This script will:
    *   Build the Go application `cmd/api-runner/main.go` into an executable named `api-runner`.
    *   Run the `api-runner` executable.

2.  **Authorization (First Run):**
    The first time you run the application, it will likely need to obtain an `access_token` and `access_secret` from Upwork.  The application will print an authorization URL to the console.
    *   Open the URL in your web browser.
    *   Authorize the application to access your Upwork account.
    *   You may be given a verifier code. Paste this code back into the console when prompted.
    *   The application will then store the `access_token` and `access_secret` in your `config.json` file for future use.

3.  **Subsequent Runs:**
    On subsequent runs, the application will use the `access_token` and `access_secret` stored in `config.json` to authenticate with the Upwork API.

## Configuration

The application's configuration is managed through:

*   **`config.json`:** Located in `internal/api_clients/upwork/config.json`. This file stores your Upwork API credentials (`consumer_key`, `consumer_secret`, `access_token`, `access_secret`).
*   **Environment Variables:** The `start.sh` script sets the `API_ENDPOINT` environment variable. You can modify this in the `start.sh` script as needed.  Other configuration parameters could be moved to environment variables for production deployments.

## License

[MIT License](LICENSE) (or specify your desired license)
