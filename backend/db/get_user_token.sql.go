// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_user_token.sql

package db

import (
	"context"
	"database/sql"
)

const getUserSessionToken = `-- name: GetUserSessionToken :one
SELECT session_token
FROM users
WHERE discord_id = $1
`

func (q *Queries) GetUserSessionToken(ctx context.Context, discordID string) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getUserSessionToken, discordID)
	var session_token sql.NullString
	err := row.Scan(&session_token)
	return session_token, err
}
