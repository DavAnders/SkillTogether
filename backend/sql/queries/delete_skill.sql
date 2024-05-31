-- name: DeleteSkill :exec
DElETE FROM skills
WHERE id = $1;