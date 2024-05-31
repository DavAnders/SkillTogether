// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: get_all_skills.sql

package db

import (
	"context"
	"database/sql"
)

const getAllSkills = `-- name: GetAllSkills :many
SELECT id, user_id, skill_description, created_at
FROM skills
`

type GetAllSkillsRow struct {
	ID               int32         `json:"id"`
	UserID           sql.NullInt32 `json:"user_id"`
	SkillDescription string        `json:"skill_description"`
	CreatedAt        sql.NullTime  `json:"created_at"`
}

func (q *Queries) GetAllSkills(ctx context.Context) ([]GetAllSkillsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllSkillsRow
	for rows.Next() {
		var i GetAllSkillsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SkillDescription,
			&i.CreatedAt,
		); err != nil {
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
