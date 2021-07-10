-- name: GetPostDiscussions :many
SELECT * FROM resource_discussions
WHERE post_id = $1
ORDER BY creation_time DESC;

-- name: GetPostDiscussionsByUser :many
SELECT * FROM resource_discussions
WHERE creator_id = $1
ORDER BY creation_time DESC;

-- name: AddPostDiscussion :one
INSERT INTO resource_discussions (
  creator_id,
  post_id,
  title,
  description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdatePostDiscussion :one
UPDATE resource_discussions
SET title = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: RemovePostDiscussion :exec
DELETE FROM resource_discussions
WHERE id = $1;

-- name: RemovePostDiscussionsByCreator :exec
DELETE FROM resource_discussions
WHERE creator_id = $1;

-- name: GetDiscussionComments :many
SELECT * FROM discussion_comments
WHERE discussion_id = $1
ORDER BY creation_time DESC;

-- name: GetDiscussionCommentsByUser :many
SELECT * FROM discussion_comments
WHERE user_id = $1
ORDER BY creation_time DESC;

-- name: UpdateDiscussionComment :one
UPDATE discussion_comments
SET content = $2
WHERE id = $1
RETURNING *;

-- name: RemoveDiscussionComment :exec
DELETE FROM discussion_comments
WHERE id = $1;

-- name: RemoveAllDiscussionComments :exec
DELETE FROM discussion_comments
WHERE discussion_id = $1;

-- name: RemoveDiscussionCommentsByUser :exec
DELETE FROM discussion_comments
WHERE discussion_id = $1 AND user_id = $2;