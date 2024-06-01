-- +goose Up
ALTER TABLE users
ADD COLUMN session_token TEXT;

-- +goose Down
ALTER TABLE users
DROP COLUMN session_token;