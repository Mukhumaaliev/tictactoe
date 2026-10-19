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
	UUID   uuid.UUID
	Board  Board
	Status GameStatus
}

type GameStatus int

const (
	GameStatusInProgress GameStatus = iota
	GameStatusDraw
	GameStatusPlayerWon
	GameStatusRobotWon
)

func NewGame() *Game {
	return &Game{
		UUID:   uuid.New(),
		Board:  Board{},
		Status: GameStatusInProgress,
	}
}
