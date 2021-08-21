package repository

import (
	"context"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
)

func (r *SqlRepository) UpdateResourceTrans(ctx context.Context, args db.UpdateResourceParams) (db.Resource, error) {
	var err error
	var updatedResource db.Resource
	err = r.execTx(ctx, func(q *db.Queries) error {
		updatedResource, err = q.UpdateResource(ctx, args)
		return err
	})
	return updatedResource, err
}
