-- name: AddUser :one
INSERT INTO users (discord_id, username, email, avatar_url, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id;