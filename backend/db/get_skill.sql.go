// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_skill.sql

package db

import (
	"context"
)

const getSkill = `-- name: GetSkill :one
SELECT id, user_id, skill_description, created_at
FROM skills
WHERE id = $1
`

func (q *Queries) GetSkill(ctx context.Context, id int32) (Skill, error) {
	row := q.db.QueryRowContext(ctx, getSkill, id)
	var i Skill
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SkillDescription,
		&i.CreatedAt,
	)
	return i, err
}
