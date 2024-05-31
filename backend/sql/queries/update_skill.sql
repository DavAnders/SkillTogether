-- name: UpdateSkill :exec
UPDATE skills
SET skill_description = $2
WHERE id = $1;