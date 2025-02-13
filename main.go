//go:generate oapi-codegen --config=config.yaml -o gen/gen.go api.yaml
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/robertjshirts/fetch-takehome/api"
	"github.com/robertjshirts/fetch-takehome/gen"
)

func main() {

	swagger, err := gen.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger: %v", err)
	}

	swagger.Servers = nil

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))

	server := api.NewReceiptHandler()
	gen.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Println("Serving on port 8080")
	log.Fatal(s.ListenAndServe())
}
