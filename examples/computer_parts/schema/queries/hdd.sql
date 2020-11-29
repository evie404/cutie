-- name: GetByID :one
SELECT * FROM hdd WHERE id = $1 LIMIT 1;

