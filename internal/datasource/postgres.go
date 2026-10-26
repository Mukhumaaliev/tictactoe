package datasource

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(lc fx.Lifecycle, cfg *Config) (*Database, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.GetDSN())
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}
	defer pool.Close()
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})
	return &Database{
		pool: pool,
	}, err
}
