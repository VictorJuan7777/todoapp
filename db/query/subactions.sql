-- name: CreateSubAction :one
INSERT INTO subactions (
    actions_id,
    title
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteSubAction :exec
DELETE FROM subactions
WHERE id = $1;

-- name: DeleteAllSubAction :exec
DELETE FROM subactions
WHERE actions_id = $1;


-- name: UpdateSubAction :one
UPDATE subactions
SET title = $2, 
completed = $3,
change_at = now()
WHERE id = $1
RETURNING *;

-- name: ListSubAction :many
SELECT * FROM subactions
WHERE actions_id = $1
ORDER BY id;