-- name: GetDiscussionCommentById :one
SELECT * FROM discussion_comments
WHERE id = $1;

-- name: GetAllDiscussionComments :many
SELECT * FROM discussion_comments
WHERE discussion_id = $1
ORDER BY id DESC LIMIT $2;

-- name: GetAllDiscussionCommentsByCursor :many
SELECT * FROM discussion_comments
WHERE discussion_id = sqlc.arg(discussionId)
AND id < sqlc.arg(cursor)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetAllUserPostComments :many
SELECT * FROM discussion_comments
WHERE discussion_id = $1 AND user_id = $2
ORDER BY creation_time DESC
LIMIT $3 OFFSET $4;

-- name: GetAllUserComments :many
SELECT * FROM discussion_comments
WHERE user_id = $1
ORDER BY creation_time DESC
LIMIT $2 OFFSET $3;

-- name: GetCommentLikes :many
SELECT * FROM comments_likes
WHERE comment_id = $1;

-- name: GetCommentLikesCount :one
SELECT COUNT(*) 
FROM comments_likes
WHERE comment_id = $1;

-- name: GetCommentDirectResponses :many
SELECT * FROM discussion_comments AS c
WHERE parent_comment_id = NULL
LIMIT $1 OFFSET $2;

-- name: AddComment :one
INSERT INTO discussion_comments (
  discussion_id, 
  parent_comment_id,
  user_id,
  content
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: RemoveComment :exec
DELETE FROM discussion_comments
WHERE id = $1;

-- name: UpdateComment :one
UPDATE discussion_comments
SET content = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: RemoveUserComments :exec
DELETE FROM discussion_comments
WHERE user_id = $1;

-- name: RemoveDiscussionComments :exec
DELETE FROM discussion_comments
WHERE discussion_id = $1;

-- name: LikeComment :exec
INSERT INTO comments_likes (
  comment_id,
  user_id
) VALUES (
  $1, $2
);

-- name: UnlikeComment :exec
DELETE FROM comments_likes
WHERE comment_id = $1
AND user_id = $2;