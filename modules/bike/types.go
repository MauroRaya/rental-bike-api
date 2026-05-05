package bike

import (
	"time"
)

type Bike struct {
	ID           int32      `json:"id"`
	Model        string     `json:"model"`
	PricePerHour int64      `json:"price_per_hour"`
	RentedAt     *time.Time `json:"rented_at"`
	RentedUntil  *time.Time `json:"rented_until"`
	CreatedAt    *time.Time `json:"created_at"`
}

type CreateBike struct {
	Model        string     `json:"model"`
	PricePerHour int64      `json:"price_per_hour"`
	RentedAt     *time.Time `json:"rented_at"`
	RentedUntil  *time.Time `json:"rented_until"`
}

type UpdateBike struct {
	Model        string     `json:"model"`
	PricePerHour int64      `json:"price_per_hour"`
	RentedAt     *time.Time `json:"rented_at"`
	RentedUntil  *time.Time `json:"rented_until"`
}
