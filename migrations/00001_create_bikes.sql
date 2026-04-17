-- +goose Up
CREATE TABLE IF NOT EXISTS bikes (
    id SERIAL PRIMARY KEY,
    model TEXT NOT NULL,
    price_per_hour BIGINT NOT NULL,
    rented_at TIMESTAMP NOT NULL,
    rented_until TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()   
);

-- +goose Down
DROP TABLE IF EXISTS bikes;
