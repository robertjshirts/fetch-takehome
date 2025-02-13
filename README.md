# Fetch Take-Home Exercise

## About
This is a take home exercise for an interview with Fetch Rewards. They supplied me an OpenAPI schema and asked for a golang or dockerized API that followed the spec. I opted to build it in golang, using Gin as my API framework (familiarity). I used oapi-codegen to generate the server framework, body validation, and path/query parameter parsing from the OpenAPI schema. I just had to implement the required methods to handle endpoints, and register them to the generated framework.

## Prerequisites
- Go 1.20+
- Git

## Setup
Clone the repo:
```bash
git clone https://github.com/robertjshirts/fetch-takehome.git
cd fetch-takehome
```
Install dependencies:
```bash
go mod tidy
```

## Running the API:
Start the server with:
`go run main.go`

## Configuration
You can change the port the server listens to by changing the `PORT` env variable in `.env`. Default is `8080`
