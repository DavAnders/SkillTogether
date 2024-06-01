-- name: GetUserByDiscordID :one
SELECT id, discord_id, username, email, avatar_url
FROM users
WHERE discord_id = $1;