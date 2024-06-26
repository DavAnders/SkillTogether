// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_interest.sql

package db

import (
	"context"
)

const getInterest = `-- name: GetInterest :one
SELECT id, user_id, interest, created_at
FROM interests
WHERE id = $1
`

func (q *Queries) GetInterest(ctx context.Context, id int32) (Interest, error) {
	row := q.db.QueryRowContext(ctx, getInterest, id)
	var i Interest
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Interest,
		&i.CreatedAt,
	)
	return i, err
}
