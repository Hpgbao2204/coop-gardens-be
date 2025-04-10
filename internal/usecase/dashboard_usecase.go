package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
)

// DashboardUsecase cung cấp các logic tổng hợp dashboard.
type DashboardUsecase struct {
	repo *repository.DashboardRepository
}

func NewDashboardUsecase(repo *repository.DashboardRepository) *DashboardUsecase {
	return &DashboardUsecase{repo: repo}
}

func (u *DashboardUsecase) GetDashboardSummary() (*models.DashboardSummary, error) {
	return u.repo.GetDashboardSummary()
}
