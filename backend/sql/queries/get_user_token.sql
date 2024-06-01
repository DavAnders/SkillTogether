-- name: GetUserSessionToken :one
SELECT session_token
FROM users
WHERE discord_id = $1;