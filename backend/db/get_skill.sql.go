// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_skill.sql

package db

import (
	"context"
	"database/sql"
)

const getSkill = `-- name: GetSkill :one
SELECT id, user_id, skill_description, created_at
FROM skills
WHERE id = $1
`

type GetSkillRow struct {
	ID               int32         `json:"id"`
	UserID           sql.NullInt32 `json:"user_id"`
	SkillDescription string        `json:"skill_description"`
	CreatedAt        sql.NullTime  `json:"created_at"`
}

func (q *Queries) GetSkill(ctx context.Context, id int32) (GetSkillRow, error) {
	row := q.db.QueryRowContext(ctx, getSkill, id)
	var i GetSkillRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SkillDescription,
		&i.CreatedAt,
	)
	return i, err
}
