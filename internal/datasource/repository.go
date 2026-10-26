package datasource

import (
	"context"
	"errors"
	"tictactoe/internal/domain"

	"github.com/jackc/pgx/v5"
)

type gameRepository struct {
	Database *Database
}

func (r *gameRepository) Save(ctx context.Context, game *domain.Game) error {
	gameModel, err := DomainToDatasource(game)
	if err != nil {
		return err
	}
	_, err = r.Database.pool.Exec(ctx, "INSERT INTO games (uuid, board, status) VALUES ($1, $2, $3) ON CONFLICT (uuid) DO UPDATE SET board = EXCLUDED.board, status = EXCLUDED.status", gameModel.UUID, gameModel.Board, gameModel.Status)
	return err
}

func (r *gameRepository) Get(ctx context.Context, uuid string) (*domain.Game, error) {

	row := r.Database.pool.QueryRow(ctx, "SELECT uuid, board, status FROM games WHERE uuid=$1", uuid)
	var game Game
	err := row.Scan(&game.UUID, &game.Board, &game.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("Игра с таким UUID не найдена")
		}
		return nil, err
	}
	return DatasourceToDomain(&game)
}

func NewGameRepository(database *Database) domain.GameRepository {
	return &gameRepository{
		Database: database,
	}
}
