package service

import (
	"math"

	"tictactoe/internal/domain"
)

// FindBestMove вычисляет наилучший ход для компьютера (CellO) алгоритмом Минимакс
func FindBestMove(board domain.Board) (domain.Move, error) {
	bestScore := math.MinInt32
	bestMove := domain.Move{Row: -1, Col: -1}

	for r := 0; r < domain.BoardSize; r++ {
		for c := 0; c < domain.BoardSize; c++ {
			if board[r][c] == domain.CellEmpty {
				board[r][c] = domain.CellO
				score := minimax(board, 0, false)
				board[r][c] = domain.CellEmpty

				if score > bestScore {
					bestScore = score
					bestMove = domain.Move{Row: r, Col: c}
				}
			}
		}
	}

	if bestMove.Row == -1 {
		return domain.Move{}, domain.ErrNoAvailableMoves
	}

	return bestMove, nil
}

func minimax(board domain.Board, depth int, isMaximizing bool) int {
	status, isOver := CheckGameOver(board)
	if isOver {
		switch status {
		case domain.StatusAIWon:
			return 10 - depth // Чем быстрее победа, тем выше оценка
		case domain.StatusPlayerWon:
			return -10 + depth // Чем позже поражение, тем выше оценка
		case domain.StatusDraw:
			return 0
		}
	}

	if isMaximizing {
		maxEval := math.MinInt32
		for r := 0; r < domain.BoardSize; r++ {
			for c := 0; c < domain.BoardSize; c++ {
				if board[r][c] == domain.CellEmpty {
					board[r][c] = domain.CellO
					eval := minimax(board, depth+1, false)
					board[r][c] = domain.CellEmpty
					if eval > maxEval {
						maxEval = eval
					}
				}
			}
		}
		return maxEval
	}

	minEval := math.MaxInt32
	for r := 0; r < domain.BoardSize; r++ {
		for c := 0; c < domain.BoardSize; c++ {
			if board[r][c] == domain.CellEmpty {
				board[r][c] = domain.CellX
				eval := minimax(board, depth+1, true)
				board[r][c] = domain.CellEmpty
				if eval < minEval {
					minEval = eval
				}
			}
		}
	}
	return minEval
}
