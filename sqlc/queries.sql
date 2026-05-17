-- name: ListBikes :many
SELECT
    id, model, price_per_hour, rented_at, rented_until, created_at
FROM
    bikes;

-- name: FindBikeByID :one
SELECT
    id, model, price_per_hour, rented_at, rented_until, created_at
FROM
    bikes
WHERE id = $1;

-- name: CreateBike :one
INSERT INTO bikes (
    model, price_per_hour, rented_at, rented_until
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;

-- name: UpdateBike :one
UPDATE bikes
SET
    model = $1,
    price_per_hour = $2,
    rented_at = $3,
    rented_until = $4
WHERE
    id = $5
RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;

-- name: DeleteBike :one
DELETE FROM
    bikes
WHERE
    id = $1
RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;

-- name: FindUserByEmail :one
SELECT
    id, email, hash
FROM
    users
WHERE
    email = $1;

-- name: CreateUser :one
INSERT INTO users (
    email, hash
) VALUES (
    $1, $2
)
RETURNING id, email;
