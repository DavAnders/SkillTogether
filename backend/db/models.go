// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Interest struct {
	ID        int32        `json:"id"`
	UserID    int32        `json:"user_id"`
	Interest  string       `json:"interest"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Skill struct {
	ID               int32         `json:"id"`
	SkillDescription string        `json:"skill_description"`
	CreatedAt        sql.NullTime  `json:"created_at"`
	UserID           sql.NullInt32 `json:"user_id"`
}

type User struct {
	ID           int32          `json:"id"`
	DiscordID    string         `json:"discord_id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	AvatarUrl    sql.NullString `json:"avatar_url"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
	SessionToken sql.NullString `json:"session_token"`
}

type UserSession struct {
	SessionToken string    `json:"session_token"`
	DiscordID    string    `json:"discord_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}
