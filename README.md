# Network Management Configuration Service

### Description
This is a Go microservice built using the Gorilla Mux framework. It is used for network management configuration and service. The service is containerized using Docker and supports Contract Testing using PACT and Load Testing using K6. It also provides documentation using Open API Swagger.

### Dependencies
This microservice has the following dependencies:

- Go 1.16 or higher
- Gorilla Mux framework
- Docker
- PACT
- K6
- Open API Swagger


### Installation

Follow the steps below to install and run the microservice:

Clone the repository using the command below:

```sh
git clone https://github.com/shubhtrino/Network-Management-Configuration-Service.git
```

Navigate to the project directory:

```sh
cd Network-Management-Configuration-Service
```

Build and run the Docker container:

```sh
docker build -t network-management-service .
docker run -p 8080:8080 network-management-service
```

Access the service using the following URL:

```sh
http://localhost:8080
```

### Usage

This microservice provides an API for managing network configurations. The available endpoints are listed below:

- GET /node - Welcome
- GET /nodes - Retrieve all configurations
- GET /node/{id} - Retrieve a configuration by ID
- POST /node - Create a new configuration
- PUT /node/{id} - Update an existing configuration
- DELETE /node/{id} - Delete a configuration by ID

### Testing
This microservice supports Contract Testing using PACT and Load Testing using K6.

Contract Testing using PACT

Load Testing using K6
K6 tests can be run using the command below:

```sh
k6 run script.js
```
Documentation
This microservice provides documentation using Open API Swagger. The Swagger UI can be accessed using the following URL:

```sh
http://localhost:8080/swagger/index.html
```


