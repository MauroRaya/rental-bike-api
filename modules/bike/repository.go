package bike

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	db *pgx.Conn
}

type Repository interface {
	ListBikes(ctx context.Context) ([]Bike, error)
	FindBikeByID(ctx context.Context, id int32) (Bike, error)
	CreateBike(ctx context.Context, arg CreateBike) (Bike, error)
	UpdateBike(ctx context.Context, arg UpdateBike, id int32) (Bike, error)
	DeleteBike(ctx context.Context, id int32) (Bike, error)
}

func NewRepository(db *pgx.Conn) Repository {
	return &repository{db}
}

func (r *repository) ListBikes(ctx context.Context) ([]Bike, error) {
	const sql = `
		SELECT
			id, model, price_per_hour, rented_at, rented_until, created_at
		FROM
			bikes;
	`

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bikes []Bike

	for rows.Next() {
		var bike Bike

		if err := rows.Scan(
			&bike.ID,
			&bike.Model,
			&bike.PricePerHour,
			&bike.RentedAt,
			&bike.RentedUntil,
			&bike.CreatedAt,
		); err != nil {
			return nil, err
		}

		bikes = append(bikes, bike)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bikes, nil
}

func (r *repository) FindBikeByID(ctx context.Context, id int32) (Bike, error) {
	const sql = `
		SELECT
			id, model, price_per_hour, rented_at, rented_until, created_at
		FROM
			bikes
		WHERE id = $1;
	`

	row := r.db.QueryRow(ctx, sql, id)

	var bike Bike

	if err := row.Scan(
		&bike.ID,
		&bike.Model,
		&bike.PricePerHour,
		&bike.RentedAt,
		&bike.RentedUntil,
		&bike.CreatedAt,
	); err != nil {
		return bike, err
	}

	return bike, nil
}

func (r *repository) CreateBike(ctx context.Context, arg CreateBike) (Bike, error) {
	const sql = `
		INSERT INTO bikes (
			model, price_per_hour, rented_at, rented_until
		) VALUES (
			$1, $2, $3, $4
		)
		RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;
	`

	row := r.db.QueryRow(ctx, sql,
		arg.Model,
		arg.PricePerHour,
		arg.RentedAt,
		arg.RentedUntil,
	)

	var bike Bike

	if err := row.Scan(
		&bike.ID,
		&bike.Model,
		&bike.PricePerHour,
		&bike.RentedAt,
		&bike.RentedUntil,
		&bike.CreatedAt,
	); err != nil {
		return bike, err
	}

	return bike, nil
}

func (r *repository) UpdateBike(ctx context.Context, arg UpdateBike, id int32) (Bike, error) {
	const sql = `
		UPDATE bikes
		SET
			model = $1,
			price_per_hour = $2,
			rented_at = $3,
			rented_until = $4
		WHERE
			id = $5
		RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;
	`

	row := r.db.QueryRow(ctx, sql,
		arg.Model,
		arg.PricePerHour,
		arg.RentedAt,
		arg.RentedUntil,
		id,
	)

	var bike Bike

	if err := row.Scan(
		&bike.ID,
		&bike.Model,
		&bike.PricePerHour,
		&bike.RentedAt,
		&bike.RentedUntil,
		&bike.CreatedAt,
	); err != nil {
		return bike, err
	}

	return bike, nil
}

func (r *repository) DeleteBike(ctx context.Context, id int32) (Bike, error) {
	const sql = `
		DELETE FROM
			bikes
		WHERE
			id = $1
		RETURNING id, model, price_per_hour, rented_at, rented_until, created_at;
	`

	row := r.db.QueryRow(ctx, sql, id)

	var bike Bike

	if err := row.Scan(
		&bike.ID,
		&bike.Model,
		&bike.PricePerHour,
		&bike.RentedAt,
		&bike.RentedUntil,
		&bike.CreatedAt,
	); err != nil {
		return bike, err
	}

	return bike, nil
}
