-- name: AddSkill :one
INSERT INTO skills (user_id, skill_description)
VALUES ($1, $2)
RETURNING id;