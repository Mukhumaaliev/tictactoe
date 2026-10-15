package datasource

// GameEntity представляет модель текущей игры для слоя хранения (datasource)
type GameEntity struct {
	ID     string       `json:"id"`
	Board  [3][3]int    `json:"board"`
	Status string       `json:"status"`
}
