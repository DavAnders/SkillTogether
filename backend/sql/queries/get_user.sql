-- name: GetUser :one
SELECT username, avatar_url 
FROM users
WHERE discord_id = $1;