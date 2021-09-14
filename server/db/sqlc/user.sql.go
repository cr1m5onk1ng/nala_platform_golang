// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addToken = `-- name: AddToken :one
INSERT INTO tokens (
    token,
    expired_at
) VALUES (
    $1, $2
) RETURNING token, created_at, expired_at, invalidated_at
`

type AddTokenParams struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (q *Queries) AddToken(ctx context.Context, arg AddTokenParams) (Token, error) {
	row := q.queryRow(ctx, q.addTokenStmt, addToken, arg.Token, arg.ExpiredAt)
	var i Token
	err := row.Scan(
		&i.Token,
		&i.CreatedAt,
		&i.ExpiredAt,
		&i.InvalidatedAt,
	)
	return i, err
}

const addUserTargetLanguage = `-- name: AddUserTargetLanguage :one
INSERT INTO learning (
    user_id,
    language
) VALUES (
    $1, $2
) RETURNING user_id, language, proficiency
`

type AddUserTargetLanguageParams struct {
	UserID   string `json:"user_id"`
	Language string `json:"language"`
}

func (q *Queries) AddUserTargetLanguage(ctx context.Context, arg AddUserTargetLanguageParams) (Learning, error) {
	row := q.queryRow(ctx, q.addUserTargetLanguageStmt, addUserTargetLanguage, arg.UserID, arg.Language)
	var i Learning
	err := row.Scan(&i.UserID, &i.Language, &i.Proficiency)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    hashed_password,
    native_language,
    role
) VALUES (
    $1, $2, $3, $4, $5, $6
) 
RETURNING id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token
`

type CreateUserParams struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	HashedPassword string         `json:"hashed_password"`
	NativeLanguage string         `json:"native_language"`
	Role           sql.NullString `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.NativeLanguage,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token FROM users
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
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.RegisteredAt,
			&i.NativeLanguage,
			&i.Role,
			&i.AccessToken,
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
SELECT id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token FROM users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}

const getUserFollowers = `-- name: GetUserFollowers :many
SELECT id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token, followed_user_id, follower_user_id, updated_at FROM users AS u
JOIN user_is_followed AS f
ON u.id = f.followed_user_id
WHERE u.id = $1 
ORDER BY u.registration_date DESC
`

type GetUserFollowersRow struct {
	ID                string         `json:"id"`
	Username          string         `json:"username"`
	Email             string         `json:"email"`
	HashedPassword    string         `json:"hashed_password"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	RegisteredAt      time.Time      `json:"registered_at"`
	NativeLanguage    string         `json:"native_language"`
	Role              sql.NullString `json:"role"`
	AccessToken       sql.NullString `json:"access_token"`
	FollowedUserID    string         `json:"followed_user_id"`
	FollowerUserID    string         `json:"follower_user_id"`
	UpdatedAt         time.Time      `json:"updated_at"`
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
			&i.ID,
			&i.Username,
			&i.Email,
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.RegisteredAt,
			&i.NativeLanguage,
			&i.Role,
			&i.AccessToken,
			&i.FollowedUserID,
			&i.FollowerUserID,
			&i.UpdatedAt,
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
SELECT user_id, language, proficiency FROM learning
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
		if err := rows.Scan(&i.UserID, &i.Language, &i.Proficiency); err != nil {
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

const invalidateToken = `-- name: InvalidateToken :exec
DELETE FROM tokens 
WHERE token = $1
`

func (q *Queries) InvalidateToken(ctx context.Context, token string) error {
	_, err := q.exec(ctx, q.invalidateTokenStmt, invalidateToken, token)
	return err
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
RETURNING id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token
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
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users 
SET role = $2
WHERE id = $1
RETURNING id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token
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
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}

const updateUserToken = `-- name: UpdateUserToken :one
UPDATE users
SET access_token = $2
WHERE email = $1
RETURNING id, username, email, hashed_password, password_changed_at, registered_at, native_language, role, access_token
`

type UpdateUserTokenParams struct {
	Email       string         `json:"email"`
	AccessToken sql.NullString `json:"access_token"`
}

func (q *Queries) UpdateUserToken(ctx context.Context, arg UpdateUserTokenParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserTokenStmt, updateUserToken, arg.Email, arg.AccessToken)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.RegisteredAt,
		&i.NativeLanguage,
		&i.Role,
		&i.AccessToken,
	)
	return i, err
}
