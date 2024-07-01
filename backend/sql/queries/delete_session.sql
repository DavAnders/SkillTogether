-- name: DeleteSession :exec
DELETE FROM user_sessions WHERE session_token = $1;