// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_user_by_discord_id.sql

package db

import (
	"context"
	"database/sql"
)

const getUserByDiscordID = `-- name: GetUserByDiscordID :one
SELECT id, discord_id, username, email, avatar_url
FROM users
WHERE discord_id = $1
`

type GetUserByDiscordIDRow struct {
	ID        int32          `json:"id"`
	DiscordID string         `json:"discord_id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	AvatarUrl sql.NullString `json:"avatar_url"`
}

func (q *Queries) GetUserByDiscordID(ctx context.Context, discordID string) (GetUserByDiscordIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByDiscordID, discordID)
	var i GetUserByDiscordIDRow
	err := row.Scan(
		&i.ID,
		&i.DiscordID,
		&i.Username,
		&i.Email,
		&i.AvatarUrl,
	)
	return i, err
}
