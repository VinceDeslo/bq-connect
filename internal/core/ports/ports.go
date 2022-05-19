package ports

import "github.com/VinceDeslo/bq-connect/internal/core/domain"

type GamesRepository interface {
	GetAll() ([]domain.Game, error)
	GetAdventure() ([]domain.Game, error)
}

type GamesService interface {
	GetAll() ([]domain.Game, error)
	GetAdventure() ([]domain.Game, error)
}
