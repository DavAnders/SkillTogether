-- name: GetAllSkills :many
SELECT id, user_id, skill_description, created_at
FROM skills;