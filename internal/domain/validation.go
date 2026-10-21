package domain

import "errors"

func UpdateGameStatus(game *Game) {
	if checkWin(game.Board, CellX) {
		game.Status = GameStatusPlayerWon
	}
	if checkWin(game.Board, CellO) {
		game.Status = GameStatusRobotWon
	}
	if isBoardFull(game.Board) {
		game.Status = GameStatusDraw
	}
}

func isValidMove(board Board, row int, col int) bool {
	if board[row][col] != CellEmpty {
		return false
	}
	return true
}

func (s *Service) IsValidBoard(oldBoard Board, newBoard Board) (bool, error) {
	if oldBoard == newBoard {
		return false, errors.New("Вы не сделали ход: доска не изменилась")
	}
	isEdited := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if oldBoard[i][j] == newBoard[i][j] {
				continue
			} else if oldBoard[i][j] != newBoard[i][j] && !isEdited {
				if !isValidMove(oldBoard, i, j) {
					return false, errors.New("Неправильный ход: клетка уже занята")
				}
				isEdited = true
			} else if oldBoard[i][j] != newBoard[i][j] && isEdited {
				if !isValidMove(oldBoard, i, j) {
					return false, errors.New("Неправильный ход: клетка уже занята")
				}
				return false, errors.New("Неправильный ход: можно изменить только одну клетку за раз")
			}
		}
	}
	return true, nil
}

func (s *Service) IsFinished(board Board) bool {
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
