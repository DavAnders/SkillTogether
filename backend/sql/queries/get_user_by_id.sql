-- name: GetUserByID :one
SELECT id, username, email, avatar_url
FROM users
WHERE id = $1;