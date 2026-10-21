package web

import (
	"tictactoe/internal/domain"
	"uuid"
)

func DomainToWeb(game *domain.Game) *Game {
	return &Game{
		UUID:   game.UUID.String(),
		Board:  game.Board,
		Status: game.Status,
	}
}

func WebToDomain(game *Game) *domain.Game {
	return &domain.Game{
		UUID:   uuid.MustParse(game.UUID),
		Board:  game.Board,
		Status: domain.GameStatus(game.Status),
	}
}
