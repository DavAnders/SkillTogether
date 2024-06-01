-- +goose Up
ALTER TABLE user_sessions ADD CONSTRAINT unique_discord_id UNIQUE (discord_id);

-- +goose Down
ALTER TABLE user_sessions DROP CONSTRAINT unique_discord_id;