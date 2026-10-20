package datasource

import "tictactoe/internal/domain"

type Game struct {
	UUID   string
	Board  domain.Board
	Status domain.GameStatus
}
