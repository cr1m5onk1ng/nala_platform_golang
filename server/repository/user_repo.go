package repository

import (
	"context"

	db "github.com/cr1m5onk1ng/nala_platform_app/db/sqlc"
)

func (r *Repository) UpdateUserLanguageTrans(ctx context.Context, params db.UpdateUserLanguageParams) (db.User, error) {
	var updatedUser db.User
	var err error
	err = r.execTx(ctx, func(q *db.Queries) error {
		updatedUser, err = q.UpdateUserLanguage(ctx, params)
		return err
	})
	return updatedUser, nil
}

func (r *Repository) UpdateUserRoleTrans(ctx context.Context, params db.UpdateUserRoleParams) (db.User, error) {
	var updatedUser db.User
	var err error
	err = r.execTx(ctx, func(q *db.Queries) error {
		updatedUser, err = q.UpdateUserRole(ctx, params)
		if err != nil {
			return err
		}
		return nil
	})
	return updatedUser, nil
}
