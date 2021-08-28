-- name: GetPostDiscussions :many
SELECT * FROM post_discussions
WHERE post_id = $1
ORDER BY creation_time DESC;

-- name: GetPostDiscussionsByUser :many
SELECT * FROM post_discussions
WHERE creator_id = $1
ORDER BY creation_time DESC;

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
