package bike

import (
	"context"
	"log/slog"

	repo "github.com/MauroRaya/bike-rental-api/sqlc"
)

type service struct {
	repo repo.Querier
}

type Service interface {
	ListBikes(ctx context.Context) ([]repo.Bike, error)
	FindBikeByID(ctx context.Context, id int32) (repo.Bike, error)
	CreateBike(ctx context.Context, arg repo.CreateBikeParams) (repo.Bike, error)
	UpdateBike(ctx context.Context, arg repo.UpdateBikeParams) (repo.Bike, error)
	DeleteBike(ctx context.Context, id int32) (repo.Bike, error)
}

func NewService(repo repo.Querier) Service {
	return &service{repo}
}

func (s *service) ListBikes(ctx context.Context) ([]repo.Bike, error) {
	bikes, err := s.repo.ListBikes(ctx)
	if err != nil {
		slog.Error("failed to list bikes", "error", err)
		return bikes, err
	}
	return bikes, nil
}

func (s *service) FindBikeByID(ctx context.Context, id int32) (repo.Bike, error) {
	bike, err := s.repo.FindBikeByID(ctx, id)
	if err != nil {
		slog.Error("failed to find bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) CreateBike(ctx context.Context, arg repo.CreateBikeParams) (repo.Bike, error) {
	bike, err := s.repo.CreateBike(ctx, arg)
	if err != nil {
		slog.Error("failed to create bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) UpdateBike(ctx context.Context, arg repo.UpdateBikeParams) (repo.Bike, error) {
	bike, err := s.repo.UpdateBike(ctx, arg)
	if err != nil {
		slog.Error("failed to update bike", "error", err)
		return bike, err
	}
	return bike, nil
}

func (s *service) DeleteBike(ctx context.Context, id int32) (repo.Bike, error) {
	bike, err := s.repo.DeleteBike(ctx, id)
	if err != nil {
		slog.Error("failed to delete bike", "error", err)
		return bike, err
	}
	return bike, nil
}
