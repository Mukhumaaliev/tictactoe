package service

import (
	"tictactoe/internal/domain"
)

// ValidateBoard проверяет корректность нового состояния поля относительно старого
func ValidateBoard(oldGame domain.Game, newBoard domain.Board) error {
	if oldGame.Status != domain.StatusInProgress {
		return domain.ErrGameOver
	}

	newPlayerMoves := 0
	newAIMoves := 0

	for r := 0; r < domain.BoardSize; r++ {
		for c := 0; c < domain.BoardSize; c++ {
			oldVal := oldGame.Board[r][c]
			newVal := newBoard[r][c]

			// Проверка допустимости значения клетки
			if newVal != domain.CellEmpty && newVal != domain.CellX && newVal != domain.CellO {
				return domain.ErrInvalidCellValue
			}

			// Проверка, что уже сделанные ходы не были изменены или стерты
			if oldVal != domain.CellEmpty {
				if newVal != oldVal {
					return domain.ErrPreviousMovesModified
				}
				continue
			}

			// Клетка была пустой в старой игре
			if newVal == domain.CellX {
				newPlayerMoves++
			} else if newVal == domain.CellO {
				newAIMoves++
			}
		}
	}

	// Компьютер не должен был делать ход от имени пользователя
	if newAIMoves > 0 {
		return domain.ErrPreviousMovesModified
	}

	// Пользователь обязан сделать ровно один ход
	if newPlayerMoves != 1 {
		return domain.ErrInvalidMoveCount
	}

	return nil
}
