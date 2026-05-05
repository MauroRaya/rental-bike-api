package bike

import (
	"context"
	"log/slog"
)

type service struct {
	repository Repository
}

type Service interface {
	ListBikes(ctx context.Context) ([]Bike, error)
	FindBikeByID(ctx context.Context, id int32) (Bike, error)
	CreateBike(ctx context.Context, arg CreateBike) (Bike, error)
	UpdateBike(ctx context.Context, arg UpdateBike, id int32) (Bike, error)
	DeleteBike(ctx context.Context, id int32) (Bike, error)
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) ListBikes(ctx context.Context) ([]Bike, error) {
	bikes, err := s.repository.ListBikes(ctx)
	if err != nil {
		slog.Error("failed to list bikes", "error", err)
		return bikes, err
	}
	return bikes, nil
}

func (s *service) FindBikeByID(ctx context.Context, id int32) (Bike, error) {
	bike, err := s.repository.FindBikeByID(ctx, id)
	if err != nil {
		slog.Error("failed to find bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) CreateBike(ctx context.Context, arg CreateBike) (Bike, error) {
	bike, err := s.repository.CreateBike(ctx, arg)
	if err != nil {
		slog.Error("failed to create bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) UpdateBike(ctx context.Context, arg UpdateBike, id int32) (Bike, error) {
	bike, err := s.repository.UpdateBike(ctx, arg, id)
	if err != nil {
		slog.Error("failed to update bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) DeleteBike(ctx context.Context, id int32) (Bike, error) {
	bike, err := s.repository.DeleteBike(ctx, id)
	if err != nil {
		slog.Error("failed to delete bike", "error", err)
		return bike, err
	}
	return bike, nil
}
