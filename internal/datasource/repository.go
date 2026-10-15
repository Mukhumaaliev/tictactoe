package datasource

import (
	"context"

	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

// Repository описывает интерфейс репозитория для слоя datasource
type Repository interface {
	Save(ctx context.Context, game domain.Game) error
	Get(ctx context.Context, id uuid.UUID) (domain.Game, error)
}
