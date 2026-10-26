package web

import "tictactoe/internal/domain"

type Game struct {
	UUID   string            `json:"uuid"`
	Board  domain.Board      `json:"board"`
	Status domain.GameStatus `json:"status"`
}
