package domain

import (
	"context"

	"github.com/google/uuid"
)

// GameRepository описывает интерфейс репозитория для работы с текущими играми
type GameRepository interface {
	// Save сохраняет текущую игру в хранилище
	Save(ctx context.Context, game Game) error

	// Get возвращает текущую игру по идентификатору
	Get(ctx context.Context, id uuid.UUID) (Game, error)
}
