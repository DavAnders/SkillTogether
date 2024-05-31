// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Skill struct {
	ID               int32         `json:"id"`
	SkillDescription string        `json:"skill_description"`
	CreatedAt        sql.NullTime  `json:"created_at"`
	UserID           sql.NullInt32 `json:"user_id"`
}

type User struct {
	ID        int32          `json:"id"`
	DiscordID string         `json:"discord_id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	AvatarUrl sql.NullString `json:"avatar_url"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}
