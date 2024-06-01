// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_user_by_id.sql

package db

import (
	"context"
	"database/sql"
)

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, avatar_url
FROM users
WHERE id = $1
`

type GetUserByIDRow struct {
	ID        int32          `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.AvatarUrl,
	)
	return i, err
}
