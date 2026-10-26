package datasource

import (
	"encoding/json"
	"tictactoe/internal/domain"
)

type Game struct {
	UUID   string            `db:"uuid"`
	Board  json.RawMessage   `db:"board"`
	Status domain.GameStatus `db:"status"`
}
