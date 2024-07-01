-- name: SimpleSearchSkill :many
SELECT skill_description, user_id, created_at
FROM skills
WHERE skill_description 
ILIKE '%' || $1 || '%';
