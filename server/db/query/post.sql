-- name: GetPosts :many
SELECT p.id AS post_id, p.user_id, p.resource_id AS resource_id, r.language, r.difficulty, r.category, r.media_type
FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.language = $1 
AND r.category = $2
AND r.media_type = $3
ORDER BY p.id LIMIT $4;

-- name: GetPostsByCursor :many
SELECT p.id AS post_id, p.user_id, p.resource_id AS resource_id, r.language, r.difficulty, r.category, r.media_type
FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE p.id < sqlc.arg(cursor)
AND r.language = sqlc.arg(language) 
AND r.category = sqlc.arg(category)
AND r.media_type = sqlc.arg(mediaType)
ORDER BY p.id 
LIMIT sqlc.arg(maxResults);

-- name: GetCommunitiesPosts :many
SELECT p.id AS post_id, p.user_id, p.resource_id AS resource_id, r.language, r.difficulty, r.category, r.media_type
FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
JOIN learning AS l
ON p.user_id = l.user_id
WHERE p.user_id = $1
AND p.language IN (
  SELECT language FROM learning AS ll
  WHERE ll.user_id = p.user_id
) 
ORDER BY p.id DESC 
LIMIT $2;

-- name: GetCommunitiesPostsByCursor :many
SELECT p.id AS post_id, p.user_id, p.resource_id AS resource_id, r.language, r.difficulty, r.category, r.media_type
FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
JOIN learning AS l
ON p.user_id = l.user_id
WHERE p.user_id = sqlc.arg(userId)
AND p.id < sqlc.arg(cursor)
AND p.language IN (
  SELECT language FROM learning AS ll
  WHERE ll.user_id = p.user_id
) 
ORDER BY p.id DESC 
LIMIT sqlc.arg(maxResults);

-- name: GetPostById :one
SELECT * FROM user_post
WHERE id = $1 LIMIT 1;

-- name: GetPostsByTopic :many
SELECT p.* FROM user_post as p
JOIN post_topics AS pt
ON p.id = pt.post_id
JOIN topics AS t
ON t.id = post_topics.topic_id
WHERE t.topic = $1
LIMIT $2 OFFSET $3;

-- name: GetPostsByLanguage :many
SELECT p.* FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.language = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3;

-- name: GetPostsByUser :many
SELECT * FROM user_post
WHERE user_id = $1
ORDER BY user_id
LIMIT $2 OFFSET $3;

-- name: GetPostsByCategory :many
SELECT p.* FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.category = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3;

-- name: GetPostsByDifficulty :many
SELECT p.* FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.difficulty = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3;

-- name: GetPostsByMediaType :many
SELECT p.* FROM user_post AS p
JOIN resources AS r
ON p.resource_id = r.id
WHERE r.media_type = $1
ORDER BY p.resource_id
LIMIT $2 OFFSET $3;

-- name: AddPost :one
INSERT INTO user_post (
  id,
  user_id, 
  resource_id, 
  post_title, 
  post_description 
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdatePost :one
UPDATE user_post 
SET post_title = $3, post_description = $4
WHERE user_id = $1 AND resource_id = $2
RETURNING *;

-- name: GetPostTags :many
SELECT t.* FROM tags as t
JOIN post_tags as pt
ON t.id = pt.tag_id
WHERE pt.post_id = $1;

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

-- name: VotePost :one
INSERT INTO votes (
  user_id,
  post_id,
  difficulty
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: RemoveVote :exec
DELETE FROM votes
WHERE user_id = $1 AND post_id = $2;

-- name: UpdateVote :one
UPDATE votes
SET difficulty = $3
WHERE user_id = $1 AND post_id = $2
RETURNING *; 

-- name: GetVote :one
SELECT * FROM votes
WHERE user_id = $1 AND post_id = $2;

-- name: GetVotes :many
SELECT * FROM votes
WHERE post_id = $1;

