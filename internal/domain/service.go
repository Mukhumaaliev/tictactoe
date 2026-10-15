package domain

import (
	"context"

	"github.com/google/uuid"
)

// GameService описывает интерфейс сервиса для работы с игрой
type GameService interface {
	// GetNextMove вычисляет следующий оптимальный ход алгоритмом «Минимакс»
	GetNextMove(game Game) (Move, error)

	// ValidateBoard проверяет, что не изменены предыдущие ходы и сделан ровно один ход
	ValidateBoard(oldGame Game, newBoard Board) error

	// CheckGameOver проверяет окончание игры и возвращает статус и флаг завершения
	CheckGameOver(board Board) (GameStatus, bool)

	// ProcessTurn обрабатывает ход игрока, делает ответный ход AI и сохраняет игру
	ProcessTurn(ctx context.Context, gameID uuid.UUID, newBoard Board) (Game, error)

	// CreateGame инициализирует новую игру с пустым полем
	CreateGame(ctx context.Context) (Game, error)

	// GetGame получает игру по идентификатору
	GetGame(ctx context.Context, gameID uuid.UUID) (Game, error)
}
