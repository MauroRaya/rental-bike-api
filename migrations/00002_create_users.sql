-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    hash TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
