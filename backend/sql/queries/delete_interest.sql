-- name: DeleteInterest :exec
DELETE FROM interests
WHERE id = $1;