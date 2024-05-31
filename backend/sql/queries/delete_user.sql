-- name: DeleteUser :exec
DELETE FROM users
WHERE discord_id = $1;