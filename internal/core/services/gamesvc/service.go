package gamesvc

import (
	"fmt"

	"github.com/VinceDeslo/bq-connect/internal/core/domain"
	"github.com/VinceDeslo/bq-connect/internal/core/ports"
)

type service struct {
	gamesRepository ports.GamesRepository
}

func New(gamesRepository ports.GamesRepository) *service {
	return &service{
		gamesRepository: gamesRepository,
	}
}

func (svc *service) GetAll() ([]domain.Game, error) {
	games, err := svc.gamesRepository.GetAll()
	if err != nil {
		return []domain.Game{}, fmt.Errorf("error fetching game from gamesRepository: %v", err)
	}
	return games, nil
}

func (svc *service) GetAdventure() ([]domain.Game, error) {
	games, err := svc.gamesRepository.GetAdventure()
	if err != nil {
		return []domain.Game{}, fmt.Errorf("error fetching game from gamesRepository: %v", err)
	}
	return games, nil
}
