-- name: AddResource :one
INSERT INTO resources (
  url, 
  language, 
  difficulty, 
  title, 
  description, 
  media_type, 
  category, 
  thumbnail_url   
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetResourceById :one
SELECT * FROM resources
WHERE id = $1;

-- name: GetResourceByUrl :one
SELECT * FROM resources
WHERE url = $1 
LIMIT 1;

-- name: GetResourcesByLanguage :many
SELECT * FROM resources
WHERE language = $1
ORDER BY inserted_at DESC;

-- name: GetResourcesPostsByUser :many
SELECT * FROM user_post
WHERE user_id = $1
ORDER BY post_time DESC;

-- name: GetResourcePost :one
SELECT * FROM user_post
WHERE resource_id = $1
LIMIT 1;

-- name: UpdateResource :one
UPDATE resources
SET url = $2, language = $3, difficulty = $4, title = $5, description = $6, media_type = $7, category = $8, thumbnail_url = $9
WHERE id = $1
RETURNING *; 

