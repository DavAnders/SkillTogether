-- name: UpdateUser :exec
UPDATE users
SET username = $2, email = $3, avatar_url = $4, updated_at = NOW()
WHERE discord_id = $1;