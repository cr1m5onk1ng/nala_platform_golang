// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    native_language,
    role
) VALUES (
    $1, $2, $3, $4, $5
) 
RETURNING id, username, email, registration_date, native_language, role
`

type CreateUserParams struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	NativeLanguage string         `json:"native_language"`
	Role           sql.NullString `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.NativeLanguage,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.RegistrationDate,
		&i.NativeLanguage,
		&i.Role,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, username, email, registration_date, native_language, role FROM users
ORDER BY id
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getAllUsersStmt, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.RegistrationDate,
			&i.NativeLanguage,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, registration_date, native_language, role FROM users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.RegistrationDate,
		&i.NativeLanguage,
		&i.Role,
	)
	return i, err
}

const getUserFollowers = `-- name: GetUserFollowers :many
SELECT follower_id, followed_id, creation_time, id, username, email, registration_date, native_language, role FROM followers AS f
JOIN users AS u
ON u.id = f.followed_id
WHERE u.id = $1 
ORDER BY u.registration_date DESC
`

type GetUserFollowersRow struct {
	FollowerID       string         `json:"follower_id"`
	FollowedID       string         `json:"followed_id"`
	CreationTime     time.Time      `json:"creation_time"`
	ID               string         `json:"id"`
	Username         string         `json:"username"`
	Email            string         `json:"email"`
	RegistrationDate time.Time      `json:"registration_date"`
	NativeLanguage   string         `json:"native_language"`
	Role             sql.NullString `json:"role"`
}

func (q *Queries) GetUserFollowers(ctx context.Context, id string) ([]GetUserFollowersRow, error) {
	rows, err := q.query(ctx, q.getUserFollowersStmt, getUserFollowers, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserFollowersRow{}
	for rows.Next() {
		var i GetUserFollowersRow
		if err := rows.Scan(
			&i.FollowerID,
			&i.FollowedID,
			&i.CreationTime,
			&i.ID,
			&i.Username,
			&i.Email,
			&i.RegistrationDate,
			&i.NativeLanguage,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTargetLanguages = `-- name: GetUserTargetLanguages :many
SELECT user_id, language FROM learning
WHERE user_id = $1
`

func (q *Queries) GetUserTargetLanguages(ctx context.Context, userID string) ([]Learning, error) {
	rows, err := q.query(ctx, q.getUserTargetLanguagesStmt, getUserTargetLanguages, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Learning{}
	for rows.Next() {
		var i Learning
		if err := rows.Scan(&i.UserID, &i.Language); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeUser = `-- name: RemoveUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) RemoveUser(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.removeUserStmt, removeUser, id)
	return err
}

const updateUserLanguage = `-- name: UpdateUserLanguage :one
UPDATE users 
SET native_language = $2
WHERE id = $1
RETURNING id, username, email, registration_date, native_language, role
`

type UpdateUserLanguageParams struct {
	ID             string `json:"id"`
	NativeLanguage string `json:"native_language"`
}

func (q *Queries) UpdateUserLanguage(ctx context.Context, arg UpdateUserLanguageParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserLanguageStmt, updateUserLanguage, arg.ID, arg.NativeLanguage)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.RegistrationDate,
		&i.NativeLanguage,
		&i.Role,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users 
SET role = $2
WHERE id = $1
RETURNING id, username, email, registration_date, native_language, role
`

type UpdateUserRoleParams struct {
	ID   string         `json:"id"`
	Role sql.NullString `json:"role"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserRoleStmt, updateUserRole, arg.ID, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.RegistrationDate,
		&i.NativeLanguage,
		&i.Role,
	)
	return i, err
}
