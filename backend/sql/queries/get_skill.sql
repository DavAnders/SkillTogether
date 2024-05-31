-- name: GetSkill :one
SELECT id, user_id, skill_description, created_at
FROM skills
WHERE id = $1;