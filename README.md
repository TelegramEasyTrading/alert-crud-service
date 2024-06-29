# Alert CRUD Application

This repository contains the source code for a CRUD (Create, Read, Update, Delete) application for managing alerts, built using Go with the Gin web framework.

## Features

- Create, read, update, and delete alerts.
- Retrieve all alerts or alerts specific to a user.

## Prerequisites

- Go 1.21.6 or higher
- Redis server for storing alert data

## Dependencies

The project uses several Go modules:

- **Gin Web Framework** for handling HTTP requests and routing.
- **GoDotEnv** for loading environment variables from a `.env` file.
- **Go-Redis** for Redis client functionality.
- **Protobuf** for data serialization.

Refer to the `go.mod` file for a detailed list of dependencies.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/TropicalDog17/alert-crud.git
   cd alert-crud
   ```

2. Install the required Go modules:
   ```bash
   go mod tidy
   ```

3. Ensure Redis server is running and accessible.

4. Create a `.env` file in the root directory with the necessary environment variables:
   ```plaintext
   REDIS_ADDR=localhost:6379
   REDIS_PASSWORD=yourpassword
   REDIS_DB=0
   ```

## Running the Application

To run the server, execute:

```bash
go run main.go
```

The server will start on `localhost:8080`.

## API Endpoints

- `GET /alert`: Retrieve a specific alert by ID.
- `POST /alert`: Create a new alert.
- `GET /alerts`: Retrieve all alerts.
- `GET /alerts/user/:userID`: Retrieve all alerts for a specific user.
- `PUT /alert`: Update an existing alert.
- `DELETE /alert`: Delete a specific alert.
- `DELETE /alerts`: Delete all alerts.

## Structure

- `main.go`: Entry point of the application, setting up the server and routes.
- `internal/`: Contains internal packages like handlers and storage logic.
- `internal/model/`: Protobuf definitions for alert data structures.

