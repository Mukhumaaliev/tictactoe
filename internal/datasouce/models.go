package datasouce

import (
	"sync"
	"tictactoe/internal/domain"
)

type Game struct {
	UUID   string
	Board  domain.Board
	Status int
}

type Storage struct {
	games sync.Map
}
