package web

import "tictactoe/internal/domain"

func DomainToWeb(game *domain.Game) *Game {
	return &Game{
		UUID:   game.UUID.String(),
		Board:  game.Board,
		Status: game.Status,
	}
}
