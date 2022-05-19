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
