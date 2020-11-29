-- name: GetCPUByID :one
SELECT * FROM cpu WHERE id = $1 LIMIT 1;

