package domain

import "errors"

var (
	ErrGameNotFound          = errors.New("game not found")
	ErrGameOver              = errors.New("game is already finished")
	ErrPreviousMovesModified = errors.New("previous moves were modified")
	ErrInvalidMoveCount      = errors.New("exactly one valid move must be made by the player")
	ErrCellNotEmpty          = errors.New("cell is already occupied")
	ErrInvalidCellValue      = errors.New("invalid cell value on the board")
	ErrNoAvailableMoves      = errors.New("no available moves left")
)
