package domain

import "uuid"

type Cell int

const (
	CellEmpty Cell = iota
	CellX
	CellO
)

type Board [3][3]Cell

type Game struct {
	uuid   uuid.UUID
	board  Board
	Status GameStatus
}

type GameStatus int

const (
	GameStatusInProgress GameStatus = iota
	GameStatusDraw
	GameStatusWon
)

func NewGame() *Game {
	return &Game{
		uuid:   uuid.New(),
		board:  Board{},
		Status: GameStatusInProgress,
	}
}
