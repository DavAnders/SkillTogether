-- name: SimpleSearchSkill :many
SELECT skill_description, user_id
FROM skills
WHERE skill_description 
ILIKE '%' || $1 || '%';
