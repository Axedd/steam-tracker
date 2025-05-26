package main

import (
	"log"
	"net/http"

	"github.com/Axedd/steam-tracker.git/internal/config"
	"github.com/Axedd/steam-tracker.git/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1) Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// 2) Open the raw *sql.DB (with error)
	sqlDB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	// Make sure to close the pool on exit
	defer sqlDB.Close()

	// 3) Wrap it in the sqlc client (no error returned)
	queries := db.New(sqlDB)

	// 4) Set up Gin and your handlers, passing `queries` where needed
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db down", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Example of listing items:
	r.GET("/items", func(c *gin.Context) {
		items, err := queries.ListItems(c) // c is a context.Context
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	// 5) Start server
	addr := ":" + cfg.HTTPPort
	log.Printf("listening on %s…", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
