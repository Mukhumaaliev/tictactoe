package datasource

import (
	"tictactoe/internal/domain"
	"uuid"
)

func DomainToDatasource(game *domain.Game) *Game {
	return &Game{
		UUID:   game.UUID.String(),
		Board:  game.Board,
		Status: game.Status,
	}
}

func DatasourceToDomain(game *Game) *domain.Game {
	return &domain.Game{
		UUID:   uuid.MustParse(game.UUID),
		Board:  game.Board,
		Status: domain.GameStatus(game.Status),
	}
}
