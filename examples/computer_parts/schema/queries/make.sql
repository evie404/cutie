-- name: GetByID :one
SELECT * FROM make WHERE id = $1 LIMIT 1;

