-- name: GetAllInterests :many
SELECT id, user_id, interest, created_at
FROM interests;