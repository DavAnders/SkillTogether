-- name: GetUserIDFromSessionToken :one
SELECT discord_id
FROM user_sessions
WHERE session_token = $1;