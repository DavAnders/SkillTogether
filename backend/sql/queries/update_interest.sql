-- name: UpdateInterest :exec
UPDATE interests
SET interest = $2
WHERE id = $1;