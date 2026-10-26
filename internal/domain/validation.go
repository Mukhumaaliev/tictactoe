package domain

import "errors"

func updateGameStatus(game *Game) {
	if checkWin(game.Board, CellX) {
		game.Status = GameStatusPlayerWon
		return
	}
	if checkWin(game.Board, CellO) {
		game.Status = GameStatusRobotWon
		return
	}
	if isBoardFull(game.Board) {
		game.Status = GameStatusDraw
	}
}

func isValidBoard(oldBoard Board, newBoard Board) error {
	var diffCount int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			oldCell := oldBoard[i][j]
			newCell := newBoard[i][j]

			if oldCell == newCell {
				continue
			}

			diffCount++
			if diffCount > 1 {
				return errors.New("неправильный ход: можно изменить только одну клетку за раз")
			}
			if oldCell != CellEmpty {
				return errors.New("неправильный ход: нельзя изменять уже занятую клетку")
			}
			if newCell != CellX {
				return errors.New("неправильный ход: игрок может ходить только крестиком (X)")
			}
		}
	}

	if diffCount == 0 {
		return errors.New("вы не сделали ход: доска не изменилась")
	}

	return nil
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
