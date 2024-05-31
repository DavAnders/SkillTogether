-- +goose Up
ALTER TABLE skills
DROP COLUMN user_id;

ALTER TABLE skills
ADD COLUMN user_id INT,
ADD FOREIGN KEY (user_id) REFERENCES users(id);

-- +goose Down
ALTER TABLE skills
DROP COLUMN user_id;

ALTER TABLE skills
ADD COLUMN user_id TEXT NOT NULL;