-- name: GetAllPostComments :many
SELECT * FROM comments
ORDER BY comment_time DESC
LIMIT $1 OFFSET $2;

-- name: GetAllUserPostComments :many
SELECT * FROM comments
WHERE post_id = $1 AND user_id = $2
ORDER BY comment_time DESC
LIMIT $3 OFFSET $4;

-- name: GetAllUserComments :many
SELECT * FROM comments
WHERE user_id = $1
ORDER BY comment_time DESC
LIMIT $2 OFFSET $3;

-- name: GetCommentLikes :many
SELECT * FROM comments_likes
WHERE comment_id = $1;

-- name: GetCommentDirectResponses :many
SELECT * FROM comments AS c
JOIN comments_responses AS cr
ON c.id = cr.source_comment_id
WHERE c.id = $1
ORDER BY c.comment_time DESC
LIMIT $2 OFFSET $3;

-- name: AddComment :one
INSERT INTO comments (
  user_id, 
  post_id,
  content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: AddCommentAsResponse :exec
INSERT INTO comments_responses (
  source_comment_id, 
  response_comment_id
) VALUES( $1, $2 );

-- name: RemoveComment :exec
DELETE FROM comments
WHERE id = $1;

-- name: RemoveUserComments :exec
DELETE FROM comments
WHERE user_id = $1;

-- name: RemovePostComments :exec
DELETE FROM comments
WHERE post_id = $1;