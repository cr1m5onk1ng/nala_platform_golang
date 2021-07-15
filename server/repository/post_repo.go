package repository

import (
	"context"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
)

func (r *Repository) UpdatePostTrans(ctx context.Context, args db.UpdatePostParams) (db.UserPost, error) {
	var updatedPost db.UserPost
	var err error
	err = r.execTx(ctx, func(q *db.Queries) error {
		updatedPost, err = q.UpdatePost(ctx, args)
		return err
	})
	return updatedPost, err
}
