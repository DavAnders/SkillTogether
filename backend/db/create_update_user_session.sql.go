// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: create_update_user_session.sql

package db

import (
	"context"
	"time"
)

const createOrUpdateUserSession = `-- name: CreateOrUpdateUserSession :exec
INSERT INTO user_sessions (discord_id, session_token, expires_at)
VALUES ($1, $2, $3)
ON CONFLICT (discord_id) DO UPDATE
SET session_token = EXCLUDED.session_token, expires_at = EXCLUDED.expires_at
`

type CreateOrUpdateUserSessionParams struct {
	DiscordID    string    `json:"discord_id"`
	SessionToken string    `json:"session_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateOrUpdateUserSession(ctx context.Context, arg CreateOrUpdateUserSessionParams) error {
	_, err := q.db.ExecContext(ctx, createOrUpdateUserSession, arg.DiscordID, arg.SessionToken, arg.ExpiresAt)
	return err
}
