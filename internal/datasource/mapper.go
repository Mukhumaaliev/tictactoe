package datasource

import (
	"github.com/google/uuid"

	"tictactoe/internal/domain"
)

// ToDomain преобразует сущность хранилища GameEntity в доменную модель Game
func ToDomain(entity GameEntity) (domain.Game, error) {
	parsedID, err := uuid.Parse(entity.ID)
	if err != nil {
		return domain.Game{}, err
	}

	return domain.Game{
		ID:     parsedID,
		Board:  domain.Board(entity.Board),
		Status: domain.GameStatus(entity.Status),
	}, nil
}

// ToEntity преобразует доменную модель Game в сущность хранилища GameEntity
func ToEntity(game domain.Game) GameEntity {
	return GameEntity{
		ID:     game.ID.String(),
		Board:  game.Board,
		Status: string(game.Status),
	}
}
