package domain

import "github.com/google/uuid"

// GameStatus представляет текущий статус игры
type GameStatus string

const (
	StatusInProgress GameStatus = "IN_PROGRESS"
	StatusPlayerWon  GameStatus = "PLAYER_WON"
	StatusAIWon      GameStatus = "AI_WON"
	StatusDraw       GameStatus = "DRAW"
)

// Game представляет модель текущей игры
type Game struct {
	ID     uuid.UUID  `json:"id"`
	Board  Board      `json:"board"`
	Status GameStatus `json:"status"`
}
