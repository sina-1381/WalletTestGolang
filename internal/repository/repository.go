package repository

import (
	"WalletTest/pkg/config"
	"WalletTest/pkg/database"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
}

type Repository struct {
	postgresql *database.Postgresql
	redis      *database.Redis
	logger     *config.Logger
}

func NewRepository(postgresql *database.Postgresql, redis *database.Redis, logger *config.Logger) *Repository {
	repository := &Repository{
		postgresql: postgresql,
		redis:      redis,
		logger:     logger,
	}
	return repository
}

func (r *Repository) WithTransaction(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := r.postgresql.Connection.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
