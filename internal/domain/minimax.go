package domain

func minimax(board Board, depth int, isMaximizing bool) int {
	if checkWin(board, CellX) {
		return -10 + depth
	}
	if checkWin(board, CellO) {
		return 10 - depth
	}
	if isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == CellEmpty {
					board[i][j] = CellO
					score := minimax(board, depth+1, false)
					board[i][j] = CellEmpty
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == CellEmpty {
					board[i][j] = CellX
					score := minimax(board, depth+1, true)
					board[i][j] = CellEmpty
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}

}

func (s *Service) FindBestMove(board Board) (int, int) {
	bestScore := -1000
	bestMove := [2]int{-1, -1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				board[i][j] = CellO
				score := minimax(board, 0, false)
				board[i][j] = CellEmpty
				if score > bestScore {
					bestScore = score
					bestMove[0] = i
					bestMove[1] = j
				}
			}
		}
	}
	return bestMove[0], bestMove[1]
}
