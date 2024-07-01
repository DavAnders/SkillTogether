-- name: SimpleSearchInterest :many
SELECT interest, user_id, created_at
FROM interests
WHERE interest
ILIKE '%' || $1 || '%';