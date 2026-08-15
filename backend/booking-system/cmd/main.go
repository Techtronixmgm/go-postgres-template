package main

import (
	"booking-system/internal/config"
	"booking-system/internal/database"
	router "booking-system/internal/routes"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config Error: %v", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("Config Validation Error: %v", err)
	}

	// Connect to PostgreSQL
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("DB Error: %v", err)
	}

	// Get Gin mode from configuration
	gin.SetMode(cfg.GinMode)

	// Create router
	r := router.NewRouter(db, cfg)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	log.Printf("Server listening on http://localhost%s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Server Failed: %v", err)
	}
}
