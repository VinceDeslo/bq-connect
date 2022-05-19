package main

import (
	"os"

	"github.com/VinceDeslo/bq-connect/internal/core/adapters/driven"
	"github.com/VinceDeslo/bq-connect/internal/core/adapters/driver"
	"github.com/VinceDeslo/bq-connect/internal/core/services/gamesvc"
	"github.com/gin-gonic/gin"
)

func main() {
	projectID := os.Getenv("BQ_CONNECT_PROJECT_ID")

	gamesRepository := driven.NewBQAdapter(projectID)
	gamesService := gamesvc.New(gamesRepository)
	gamesHandler := driver.NewHTTPAdapter(gamesService)

	router := gin.New()
	router.GET("/games", gamesHandler.GetAll)
	router.GET("/games/adventure", gamesHandler.GetAdventure)
	router.Run(":8080")

	gamesRepository.Close()
}
