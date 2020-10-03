-- name: GetTwitterUserByID :one
SELECT * FROM twitter_users WHERE id = $1 LIMIT 1;

-- name: GetTwitterUserByScreenName :one
SELECT * FROM twitter_users WHERE screen_name = $1 LIMIT 1;

-- name: GetScreenNameByID :one
SELECT screen_name FROM twitter_users WHERE id = $1 LIMIT 1;

-- name: GetIDByScreenName :one
SELECT id FROM twitter_users WHERE screen_name = $1 LIMIT 1;
