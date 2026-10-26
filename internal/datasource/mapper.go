package datasource

import (
	"encoding/json"
	"tictactoe/internal/domain"
	"uuid"
)

func DomainToDatasource(game *domain.Game) (*Game, error) {
	board, err := json.Marshal(game.Board)
	if err != nil {
		return nil, err
	}
	return &Game{
		UUID:   game.UUID.String(),
		Board:  board,
		Status: game.Status,
	}, nil
}

func DatasourceToDomain(game *Game) (*domain.Game, error) {
	var board domain.Board
	if err := json.Unmarshal(game.Board, &board); err != nil {
		return nil, err
	}
	return &domain.Game{
		UUID:   uuid.MustParse(game.UUID),
		Board:  board,
		Status: domain.GameStatus(game.Status),
	}, nil
}
