-- name: CreateOrUpdateUserSession :exec
INSERT INTO user_sessions (discord_id, session_token, expires_at)
VALUES ($1, $2, $3)
ON CONFLICT (discord_id) DO UPDATE
SET session_token = EXCLUDED.session_token, expires_at = EXCLUDED.expires_at;