//go:generate oapi-codegen --config=config.yaml -o gen/gen.go api.yaml
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/robertjshirts/fetch-takehome/api"
	"github.com/robertjshirts/fetch-takehome/gen"
)

func main() {
	godotenv.Load()

	port := getEnv("PORT", "8080")

	router := setupRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	go func() {
		log.Printf("Server running on port %s", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	gracefulShutdown(server)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func setupRouter() *gin.Engine {
	swagger, err := gen.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger: %v", err)
	}

	swagger.Servers = nil

	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))

	server := api.NewReceiptHandler()
	gen.RegisterHandlers(router, server)

	return router
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully shutdown")
}
