-- name: GetUserByID :one
SELECT id, username, email, avatar_url, discord_id
FROM users
WHERE id = $1;