-- name: AddInterest :one
INSERT INTO interests (user_id, interest)
VALUES ($1, $2)
RETURNING id;