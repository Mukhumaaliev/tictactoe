package domain

type Service struct {
}

type GameService interface {
	FindBestMove(board Board) (int, int)
	IsValidBoard(oldBoard Board, newBoard Board) (bool, error)
	IsFinished(board Board) bool
}
