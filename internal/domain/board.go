package domain

const BoardSize = 3

// Значения клеток на игровом поле
const (
	CellEmpty = 0 // Пустая клетка
	CellX     = 1 // Ход игрока (крестик)
	CellO     = 2 // Ход компьютера (нолик)
)

// Board представляет собой игровое поле 3x3 в виде целочисленной матрицы
type Board [BoardSize][BoardSize]int

// Move представляет координаты хода
type Move struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// Clone создает глубокую копию игрового поля
func (b Board) Clone() Board {
	var clone Board
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			clone[r][c] = b[r][c]
		}
	}
	return clone
}

// IsFull проверяет, заполнено ли все поле
func (b Board) IsFull() bool {
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			if b[r][c] == CellEmpty {
				return false
			}
		}
	}
	return true
}
