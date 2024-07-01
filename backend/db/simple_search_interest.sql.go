// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: simple_search_interest.sql

package db

import (
	"context"
	"database/sql"
)

const simpleSearchInterest = `-- name: SimpleSearchInterest :many
SELECT interest, user_id, created_at
FROM interests
WHERE interest
ILIKE '%' || $1 || '%'
`

type SimpleSearchInterestRow struct {
	Interest  string       `json:"interest"`
	UserID    int32        `json:"user_id"`
	CreatedAt sql.NullTime `json:"created_at"`
}

func (q *Queries) SimpleSearchInterest(ctx context.Context, dollar_1 sql.NullString) ([]SimpleSearchInterestRow, error) {
	rows, err := q.db.QueryContext(ctx, simpleSearchInterest, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SimpleSearchInterestRow
	for rows.Next() {
		var i SimpleSearchInterestRow
		if err := rows.Scan(&i.Interest, &i.UserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
