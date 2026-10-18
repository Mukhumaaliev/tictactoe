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

func IsValidMove(board Board, row int, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false
	}
	if board[row][col] != CellEmpty {
		return false
	}
	return true
}

func FindBestMove(board Board) (int, int) {
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

func isFinished(board Board) bool {
	if checkWin(board, CellX) || checkWin(board, CellO) || isBoardFull(board) {
		return true
	}
	return false
}

func checkWin(board Board, player Cell) bool {
	if board[0][0] == player && board[1][1] == player && board[2][2] == player || board[0][2] == player && board[1][1] == player && board[2][0] == player {

		return true
	}
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player || board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	return false
}

func isBoardFull(board Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == CellEmpty {
				return false
			}
		}
	}
	return true
}
