-- name: GetPostById :one
SELECT * FROM user_post
WHERE id = $1;

-- name: GetPostsByLanguage :many
SELECT * FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.language = $1;

-- name: GetPostsByUser :many
SELECT * FROM user_post
WHERE user_id = $1;

-- name: GetPostsByCategory :many
SELECT * FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.category = $1;

-- name: GetPostsByDifficulty :many
SELECT * FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.difficulty = $1;

-- name: GetPostsByMediaType :many
SELECT * FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.media_type = $1;

-- name: AddPost :one
INSERT INTO user_post (
  user_id, 
  resource_id, 
  post_title, 
  post_description 
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdatePost :one
UPDATE user_post 
SET post_title = $3, post_description = $4
WHERE user_id = $1 AND resource_id = $2
RETURNING *;

-- name: GetPostTags :many
SELECT * FROM tags
WHERE post_id = $1;

-- name: GetPostDifficultyVotes :many
SELECT * from votes
WHERE post_id = $1;

-- name: GetPostLikes :many
SELECT * from likes
WHERE post_id = $1;

-- name: RemovePost :exec
DELETE FROM user_post 
WHERE id = $1;

-- name: RemoveUserPosts :exec
DELETE FROM user_post
WHERE user_id = $1;

