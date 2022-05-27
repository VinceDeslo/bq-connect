package main

import (
	"log"
	"os"

	"github.com/VinceDeslo/bq-connect/internal/core/adapters/driven"
	"github.com/VinceDeslo/bq-connect/internal/core/adapters/driver"
	"github.com/VinceDeslo/bq-connect/internal/core/services/gamesvc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Loading environment variables failed: %v", err)
	}
	projectID := os.Getenv("BQ_CONNECT_PROJECT_ID")

	gamesRepository := driven.NewBQAdapter(projectID)
	gamesService := gamesvc.New(gamesRepository)
	gamesHandler := driver.NewHTTPAdapter(gamesService)

	router := gin.New()
	router.Use(driver.CorsMiddleware())
	router.GET("/games", gamesHandler.GetAll)
	router.GET("/games/adventure", gamesHandler.GetAdventure)
	router.Run(":8080")

	gamesRepository.Close()
}
