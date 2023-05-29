-- name: CreateAction :one
INSERT INTO actions (
    title,
    username
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeleteAction :exec
DELETE FROM actions
WHERE id =$1;

-- name: UpdateAction :one
UPDATE actions
SET completed = $2,
    title = $3,
    change_at = now()
WHERE id = $1
RETURNING *;

-- name: ListAction :many
SELECT * FROM actions
WHERE username = $1
ORDER BY id;