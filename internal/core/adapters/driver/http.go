package driver

import (
	"github.com/VinceDeslo/bq-connect/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HTTPAdapter struct {
	service ports.GamesService
}

func NewHTTPAdapter(service ports.GamesService) *HTTPAdapter {
	return &HTTPAdapter{
		service: service,
	}
}

func (adapter *HTTPAdapter) GetAll(ctx *gin.Context) {
	games, err := adapter.service.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, games)
}

func (adapter *HTTPAdapter) GetAdventure(ctx *gin.Context) {
	games, err := adapter.service.GetAdventure()
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, games)
}

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}
}
