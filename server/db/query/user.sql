-- name: CreateUser :one
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
RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY id;

-- name: GetUserFollowers :many
SELECT * FROM followers AS f
JOIN users AS u
ON u.id = f.followed_id
WHERE u.id = $1 
ORDER BY u.registration_date DESC;

-- name: AddUserTargetLanguage :one
INSERT INTO learning (
    user_id,
    language
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateUserLanguage :one
UPDATE users 
SET native_language = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserRole :one
UPDATE users 
SET role = $2
WHERE id = $1
RETURNING *;

-- name: RemoveUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserTargetLanguages :many
SELECT * FROM learning
WHERE user_id = $1;
