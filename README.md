# Orders API `Microservice Architecture`

## Description

An Orders API microservice architecture is a design approach where the functionality related to managing orders within a system is broken down into smaller, specialized services that communicate with each other via APIs (Application Programming Interfaces).

### Microservices

The system is composed of multiple small, independent services, each responsible for a specific aspect of order management such as order creation, processing, fulfillment, and tracking.

## Installation

To Install the project, follow these steps:

1. Clone the repository: `git clone https://github.com/pranshu321/golang-orders-api.git`
2. Install the necessary dependencies
3. Configure the environment variables: Create a `.env` file and set the required variables.
4. Build the project: `go build`
5. Start the server: `go run main.go`

## Routes

These are the available routes in API. Includes the HTTP methods, endpoints, and any required parameters or headers.

- `GET /orders`: Retrieve all orders.
- `POST /orders`: Create a new order.
- `GET /orders/{id}`: Retrieve a specific order by ID.
- `PUT /orders/{id}`: Update an existing order.
- `DELETE /orders/{id}`: Delete an order.

## Technologies Used

This project utilizes the following technologies and frameworks:

- go
- chi package
- mongoDb

## Features

The main features of this project include:

- Create new orders
- Retrieve all orders
- Retrieve a specific order by ID
- Update an existing order
- Delete an order

These features provide users with the ability to manage orders efficiently.

## Contact

- [Email](mailto:pranshujain0111@gmail.com)
