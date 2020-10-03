-- name: GetTweetByID :one
SELECT * FROM twitter_tweets WHERE id = $1 LIMIT 1;

-- name: GetTweetsByUserID :many
SELECT * FROM twitter_tweets WHERE user_id = $1;
