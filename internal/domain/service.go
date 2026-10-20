package domain

import "uuid"

type Service struct {
	repository GameRepository
}

type GameService interface {
	FindBestMove(board Board) (int, int)
	IsValidBoard(oldBoard Board, newBoard Board) (bool, error)
	IsFinished(board Board) bool
}

type GameRepository interface {
	Save(g Game)
	Get(uuid uuid.UUID) (Game, error)
}
