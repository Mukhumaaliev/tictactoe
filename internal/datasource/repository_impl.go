package datasource

import (
	"context"

	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

type gameRepository struct {
	storage *Storage
}

// NewGameRepository создает реализацию domain.GameRepository на основе Storage
func NewGameRepository(storage *Storage) domain.GameRepository {
	return &gameRepository{
		storage: storage,
	}
}

// Save сохраняет доменную модель игры в хранилище через маппер
func (r *gameRepository) Save(ctx context.Context, game domain.Game) error {
	entity := ToEntity(game)
	r.storage.Set(entity.ID, entity)
	return nil
}

// Get возвращает доменную модель игры из хранилища по ID
func (r *gameRepository) Get(ctx context.Context, id uuid.UUID) (domain.Game, error) {
	entity, ok := r.storage.Get(id.String())
	if !ok {
		return domain.Game{}, domain.ErrGameNotFound
	}

	game, err := ToDomain(entity)
	if err != nil {
		return domain.Game{}, err
	}

	return game, nil
}
