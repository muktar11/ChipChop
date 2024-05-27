package service

import (
    "context"
    "github.com/your_project/internal/domain"
    "github.com/your_project/internal/repository"
)

type EstimationService struct {
    repo repository.EstimationRepository
}

func NewEstimationService(repo repository.EstimationRepository) *EstimationService {
    return &EstimationService{repo: repo}
}

func (s *EstimationService) CreateEstimation(ctx context.Context, estimation domain.Estimation) error {
    return s.repo.CreateEstimation(ctx, estimation)
}

func (s *EstimationService) GetEstimationByID(ctx context.Context, id int64) (domain.Estimation, error) {
    return s.repo.GetEstimationByID(ctx, id)
}

func (s *EstimationService) ListEstimations(ctx context.Context, limit, offset int32) ([]domain.Estimation, error) {
    return s.repo.ListEstimations(ctx, limit, offset)
}

func (s *EstimationService) UpdateEstimation(ctx context.Context, id int64, price *float64) error {
    return s.repo.UpdateEstimation(ctx, id, price)
}

func (s *EstimationService) DeleteEstimation(ctx context.Context, id int64) error {
    return s.repo.DeleteEstimation(ctx, id)
}
