-- +goose Up
CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    skill_description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE skills;