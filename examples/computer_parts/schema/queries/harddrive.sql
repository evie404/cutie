-- name: GetHDDByID :one
SELECT * FROM harddrives WHERE id = $1 LIMIT 1;

