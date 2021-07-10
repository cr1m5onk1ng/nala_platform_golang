package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	*Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:      db,
		Queries: New(db),
	}
}

func (Repository *Repository) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := Repository.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
