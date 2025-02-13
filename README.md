# Fetch Take-Home Exercise

## About
This is a take home exercise for an interview with Fetch Rewards. They supplied me an OpenAPI schema and asked for a golang or dockerized API that followed the spec. I opted to build it in golang. I used a tool called [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) with the [Gin API framework](https://github.com/gin-gonic/gin) to generate almost all of the boilerplate, validation included. Pretty neat tool.

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
