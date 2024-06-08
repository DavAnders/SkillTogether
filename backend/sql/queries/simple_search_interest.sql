-- name: SimpleSearchInterest :many
SELECT interest, user_id
FROM interests
WHERE interest
ILIKE '%' || $1 || '%';