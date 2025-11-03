# Go API Service

A simple API service built with Go and the [Gin](https://gin-gonic.com/) framework. This service provides two sets of endpoints:

1.  A default set of RESTful endpoints for typical CRUD-like operations.
2.  An inbound endpoint designed to receive and log requests from third-party systems, which is useful for debugging and integration testing.

## Features

-   RESTful API endpoints (`/api/default`).
-   Inbound webhook catcher to inspect raw requests from external services (`/api/inbound`).
-   Configurable port via environment variables.
-   Cross-Origin Resource Sharing (CORS) middleware enabled to allow requests from any origin.

## Prerequisites

-   Go (version 1.18 or higher is recommended).

## Getting Started

Follow these instructions to get the project up and running on your local machine.

### 1. Clone the repository

```sh
git clone <your-repository-url>
cd api-service-go
```

### 2. Install dependencies

This project uses Go Modules. To install the necessary dependencies, run:

```sh
go mod tidy
```

### 3. Configuration

The application can be configured using a `.env` file in the root of the project.

Create a file named `.env`:

```sh
touch .env
```

Add the following environment variable to it. If this file or variable is not present, the service will default to port `8080`.

```env
# .env
API_SERVICE_PORT=8080
```

## Running the Application

You can run the application in two ways:

### Method 1: Using `go run`

This method is ideal for development as it compiles and runs the application in one step.

```sh
go run main.go
```

### Method 2: Building and running an executable file

This method compiles the application into a single executable file (`.exe` on Windows, or a binary on macOS/Linux). This is the standard way to deploy Go applications.

1.  **Build the executable:**

    ```sh
    # For macOS/Linux
    go build -o api-service main.go

    # For Windows
    go build -o api-service.exe main.go
    ```

2.  **Run the executable:**

    ```sh
    # On macOS/Linux
    ./api-service

    # On Windows
    .\api-service.exe
    ```

After starting, the server will be listening on the configured port (e.g., `http://localhost:8080`).

## API Endpoints

The API is available under the `/api` prefix.

### Default API (`/api/default`)

This is a standard set of REST endpoints.

-   `GET /api/default`: Get all items.
-   `GET /api/default/:id`: Get an item by ID.
-   `POST /api/default`: Create a new item.
-   `PUT /api/default/:id`: Update an item by ID.
-   `DELETE /api/default/:id`: Delete an item by ID.

### Inbound Webhook API (`/api/inbound`)

These endpoints are designed to catch and log any incoming request (headers, query parameters, and body) to the console. This is very useful when you need to inspect the exact data a third-party service is sending to your webhook.

You can point a third-party service to `http://<your-server-ip>:<port>/api/inbound`.

-   `GET /api/inbound`
-   `POST /api/inbound`
-   `PUT /api/inbound`
-   `DELETE /api/inbound`

These endpoints also accept an optional `:id` parameter (e.g., `/api/inbound/123`).

This `README.md` file provides clear instructions on how to set up, configure, and run your project, including how to create and use an executable file as you requested. Let me know if you have any other questions!

<!--
[PROMPT_SUGGESTION]Can you add some `curl` examples to the README to show how to test the API endpoints?[/PROMPT_SUGGESTION]
[PROMPT_SUGGESTION]Refactor the `inbound_handler.go` to reduce code duplication across the different HTTP methods.[/PROMPT_SUGGESTION]
