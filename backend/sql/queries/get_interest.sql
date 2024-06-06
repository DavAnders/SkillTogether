-- name: GetInterest :one
SELECT id, user_id, interest, created_at
FROM interests
WHERE id = $1;