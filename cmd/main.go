package main

import (
	"log"
	"net/http"
	"github.com/Abhilash3108/E-commerce-with-go/internal/api"
	"github.com/Abhilash3108/E-commerce-with-go/pkg/config"
	"github.com/Abhilash3108/E-commerce-with-go/pkg/logger"
)


func main() {
    cfg := config.LoadConfig()           // Load configurations
    logger.InitLogger(cfg.LogLevel)      // Initialize logger

    router := api.NewRouter()             // Setup routes

    log.Println("Starting server on :8000")
    if err := http.ListenAndServe(":8000", router); err != nil {
        log.Fatal(err)
    }
}