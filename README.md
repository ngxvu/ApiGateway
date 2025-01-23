![img_1.png](img_1.png)

Simple API Gateway For Route Api From Third Party Service
==================================
## Description
This is a simple API Gateway for routing API requests from a third-party service to a local service.
It also converts the third-party API to a new API endpoint for the local service so that the local service can public API endpoints.
The API Gateway is built using Golang Language and uses a PostgreSQL database to store logs of the API requests.

## Features
- Client: Sends HTTP requests to the API Gateway.
- API Gateway: Routes requests to either the third-party API or the local service.
- Third-Party API: External service that the API Gateway interacts with (e.g., Goong API).
- Database: PostgreSQL database used to store logs and other data.

## Project Layout
```
3rd-party-gateway/
├── config/
│   ├── postgres.go
│   ├── api_key.go
├── database/
│   ├── docker_compose.yml
│   ├── logs.sql
├── middleware/
│   ├── proxies.go
├── models/
│   ├── models.go
├── services/
│   ├── services_a.go
├── .env
├── .gitignore
├── go.mod
├── main.go
├── README.md
```

## Project Layout
- `config/`: Contains configuration files for the project.
- `config/postgres.go`: Contains configuration for the PostgreSQL database.
- `config/api_key.go`: Contains API key configuration for the project.
- `database/`: Contains database configuration files.
- `database/docker_compose.yml`: Contains Docker Compose configuration for the PostgreSQL database.
- `database/logs.sql`: Contains SQL queries for creating tables in the PostgreSQL database.
- `middleware/`: Contains middleware functions for the project.
- `middleware/proxies.go`: Contains proxy middleware functions for the project.
- `models/`: Contains models for the project.
- `models/models.go`: Contains models for the project.
- `services/`: Contains services for the project.
- `services/services_a.go`: Contains service_a functions for the project.
- `.env`: Contains environment variables for the project.
- `.gitignore`: Contains files and directories to ignore.
- `go.mod`: Contains dependencies for the project.
- `main.go`: Contains the main function for the project.
- `README.md`: Contains information about the project.

## Requirement
- Requires [Go](https://golang.org/dl/) v1.23+ to run.
- Requires [Docker](https://docs.docker.com/get-docker/) v20.10.7+ to run.
- Requires [Docker Compose](https://docs.docker.com/compose/install/) v1.29.2+ to run.
- Requires [PostgreSQL](https://www.postgresql.org/download/) v13.3+ to run.
- Requires [Postman](https://www.postman.com/downloads/) v8.10.0+ to run.

## Installation
1. Clone the repository:
   ```bash
   git clone -b medium git@github.com:ngxvu/3rd-party-gateway
    ```
2. Change into the project directory:
   ```bash
   cd 3rd-party-gateway
   ```

3. Create a `.env` file in the project root directory and add the following environment variables:
   ```bash
    DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=api_gateway
   DB_HOST=localhost
   DB_PORT=5432
   PORT=8081
    ```
4. Start the PostgreSQL database using Docker Compose:  
   ```bash
   docker-compose -f database/docker_compose.yml up -d
    ```
5. Create the tables in the PostgreSQL database:
   ```bash
   psql -h localhost -U postgres -d api_gateway -a -f database/logs.sql
    ```
6. Run the following command to start the API Gateway:
   ```bash
   go run main.go
   ```
The API Gateway will start running on http://localhost:8081.  
   7. You can now make requests to the API Gateway using the following endpoints:  
   - GET /api/v1/service_a/endpoint-1
   - GET /api/v1/service_a/endpoint-2
   
This is just an example of how to use the API Gateway. You can add more services and endpoints as needed.# 3rd-party-gateway
