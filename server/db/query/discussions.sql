-- name: GetPostDiscussions :many
SELECT * FROM post_discussions
WHERE post_id = sqlc.arg(postId)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetPostDiscussionsByCursor :many
SELECT * FROM post_discussions
WHERE post_id = sqlc.arg(postId)
AND id < sqlc.arg(cursor)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetPostDiscussionsByUser :many
SELECT * FROM post_discussions
WHERE creator_id = sqlc.arg(creatorId)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetPostDiscussionsByUserByCursor :many
SELECT * FROM post_discussions
WHERE creator_id = sqlc.arg(creatorId)
AND id < slqc.arg(cursor)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetPostDiscussionById :one
SELECT * FROM post_discussions
WHERE id = $1;

-- name: AddPostDiscussion :one
INSERT INTO post_discussions (
  creator_id,
  post_id,
  title,
  description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdatePostDiscussion :one
UPDATE post_discussions
SET title = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: RemovePostDiscussion :exec
DELETE FROM post_discussions
WHERE id = $1;

-- name: RemovePostDiscussionsByCreator :exec
DELETE FROM post_discussions
WHERE creator_id = $1;
