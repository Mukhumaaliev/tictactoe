package domain

type GameService interface {
	FindBestMove(board Board) (int, int)
	IsValidMove(board Board, row int, col int) bool
	IsFinished(board Board) bool
}
