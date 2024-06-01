-- +goose Up
ALTER TABLE user_sessions DROP CONSTRAINT IF EXISTS user_sessions_user_id_fkey;

ALTER TABLE user_sessions ALTER COLUMN user_id TYPE TEXT;

ALTER TABLE user_sessions RENAME COLUMN user_id TO discord_id;

ALTER TABLE user_sessions ADD CONSTRAINT user_sessions_discord_id_fkey FOREIGN KEY (discord_id) REFERENCES users(discord_id);

-- +goose Down
ALTER TABLE user_sessions DROP CONSTRAINT IF EXISTS user_sessions_discord_id_fkey;

ALTER TABLE user_sessions RENAME COLUMN discord_id TO user_id;

ALTER TABLE user_sessions ALTER COLUMN user_id TYPE INTEGER;

ALTER TABLE user_sessions ADD CONSTRAINT user_sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);