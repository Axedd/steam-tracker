// cmd/server/main.go
package main

import (
	"log"

	"github.com/Axedd/steam-tracker.git/internal/api"
	"github.com/Axedd/steam-tracker.git/internal/config"
	"github.com/Axedd/steam-tracker.git/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect DB
	sqlDB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer sqlDB.Close()

	queries := db.New(sqlDB)

	r := gin.Default()

	// ─── Enable CORS ─────────────────────────────────────────────────
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // your Vite dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// ────────────────────────────────────────────────────────────────

	// Your existing routes
	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		if err := sqlDB.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "db down", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Mount your appids at the root
	appHandler := api.NewAppIDHandler(queries)
	appHandler.RegisterRoutes(r)

	paramsHandler := api.NewSteamParamHandler(queries)
	paramsHandler.RegisterRoutes(r)

	addr := ":" + cfg.HTTPPort
	log.Printf("listening on %s…", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
