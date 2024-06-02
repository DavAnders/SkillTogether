-- name: GetUserIDByDiscordID :one
SELECT id
FROM users
WHERE discord_id = $1;