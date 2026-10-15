package service

import "tictactoe/internal/domain"

// CheckGameOver проверяет, завершилась ли игра, и возвращает статус игры и флаг окончания
func CheckGameOver(board domain.Board) (domain.GameStatus, bool) {
	// Проверка строк
	for r := 0; r < domain.BoardSize; r++ {
		if board[r][0] != domain.CellEmpty &&
			board[r][0] == board[r][1] &&
			board[r][1] == board[r][2] {
			return winnerStatus(board[r][0]), true
		}
	}

	// Проверка столбцов
	for c := 0; c < domain.BoardSize; c++ {
		if board[0][c] != domain.CellEmpty &&
			board[0][c] == board[1][c] &&
			board[1][c] == board[2][c] {
			return winnerStatus(board[0][c]), true
		}
	}

	// Проверка главной диагонали
	if board[0][0] != domain.CellEmpty &&
		board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] {
		return winnerStatus(board[0][0]), true
	}

	// Проверка побочной диагонали
	if board[0][2] != domain.CellEmpty &&
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] {
		return winnerStatus(board[0][2]), true
	}

	// Проверка на ничью (если все клетки заполнены)
	if board.IsFull() {
		return domain.StatusDraw, true
	}

	return domain.StatusInProgress, false
}

func winnerStatus(cell int) domain.GameStatus {
	if cell == domain.CellX {
		return domain.StatusPlayerWon
	}
	return domain.StatusAIWon
}
