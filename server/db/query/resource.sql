-- name: AddResource :one
INSERT INTO resources (
  url, 
  language, 
  difficulty,
  media_type, 
  category
) VALUES (
    $1, $2, $3, $4, $5
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
ORDER BY id DESC LIMIT $2;

-- name: GetResourcesByLanguageByCursor :many
SELECT * FROM resources
WHERE language = sqlc.arg(language)
AND id < sqlc.arg(cursor)
ORDER BY id DESC LIMIT sqlc.arg(maxResults);

-- name: GetResourcesPostsByUser :many
SELECT r.* FROM resources AS r
JOIN user_post AS p
ON r.id = p.resource_id
WHERE p.user_id = $1
ORDER BY id DESC LIMIT $2;

-- name: GetResourcePost :one
SELECT * FROM user_post
WHERE resource_id = $1
LIMIT 1;

-- name: UpdateResource :one
UPDATE resources
SET url = $2, language = $3, difficulty = $4, media_type = $5, category = $6
WHERE id = $1
RETURNING *; 

