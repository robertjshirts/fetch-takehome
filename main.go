//go:generate oapi-codegen --config=config.yaml -o gen/api.gen.go api.yaml
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertjshirts/fetch-takehome/api"
	"github.com/robertjshirts/fetch-takehome/gen"
)

func main() {
	server := api.NewReceiptHandler()

	r := gin.Default()

	gen.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
