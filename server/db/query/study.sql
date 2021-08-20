-- name: GetUserStudyLists :many
SELECT * FROM study_lists
WHERE user_id = $1
ORDER BY creation_time DESC;

-- name: GetStudyList :one
SELECT * FROM study_lists
WHERE id = $1;

-- name: AddStudyList :one
INSERT INTO study_lists (
  user_id,
  title,
  description 
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateStudyList :one
UPDATE study_lists
SET title = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: RemoveStudyList :exec
DELETE FROM study_lists
WHERE id = $1;

-- name: RemoveAllUserStudyLists :exec
DELETE FROM study_lists
WHERE user_id = $1;

-- name: GetStudyListResources :many
SELECT r.* FROM study_list_resource AS slr
JOIN resources r
ON r.id = slr.resource_id
WHERE slr.study_list_id = $1
ORDER BY slr.time_added DESC;

-- name: GetUserSavedResources :many
SELECT * FROM study_list_resource AS slr
JOIN resources AS r
ON r.id = slr.resource_id
JOIN study_lists AS sl
ON sl.id = slr.study_list_id
WHERE sl.user_id = $1
ORDER BY slr.time_added DESC;

-- name: AddResourceToStudyList :one
INSERT INTO study_list_resource (
  study_list_id,
  resource_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: RemoveResourceFromStudyList :exec
DELETE FROM study_list_resource
WHERE resource_id = $1;
