package repository

import (
	"context"
	"database/sql"
	"fmt"

	sqlc "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
)

type Repository struct {
	*sqlc.Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (r *Repository) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
