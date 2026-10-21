package datasource

import (
	"tictactoe/internal/domain"
	"uuid"
)

type gameRepository struct {
	storage *Storage
}

func (r *gameRepository) Save(g *domain.Game) {
	game := DomainToDatasource(g)
	r.storage.saveGame(game)
}

func (r *gameRepository) Get(uuid uuid.UUID) (*domain.Game, error) {
	game, err := r.storage.getGame(uuid.String())
	if err != nil {
		return &domain.Game{}, err
	}
	return DatasourceToDomain(game), nil
}

func NewGameRepository(storage *Storage) domain.GameRepository {
	return &gameRepository{
		storage: storage,
	}
}
