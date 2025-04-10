package repository

import (
    "coop-gardens-be/internal/models"

    "gorm.io/gorm"
)

type DashboardRepository struct {
    DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
    return &DashboardRepository{DB: db}
}

func (r *DashboardRepository) GetDashboardSummary() (*models.DashboardSummary, error) {
    summary := &models.DashboardSummary{}

    if err := r.DB.Raw("SELECT COUNT(*) FROM crops").Scan(&summary.TotalCrops).Error; err != nil {
        return nil, err
    }

    if err := r.DB.Raw("SELECT COUNT(*) FROM seasons").Scan(&summary.TotalSeasons).Error; err != nil {
        return nil, err
    }

    if err := r.DB.Raw("SELECT COALESCE(AVG(rating), 0) FROM reviews").Scan(&summary.AvgProductRating).Error; err != nil {
        return nil, err
    }

    if err := r.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&summary.TotalOrders).Error; err != nil {
        return nil, err
    }

    type RoleCount struct {
        Role  string
        Count int
    }
    var roleCounts []RoleCount
    if err := r.DB.Raw(`
        SELECT r.name as role, COUNT(*) as count
        FROM user_roles ur JOIN roles r ON ur.role_id = r.id 
        GROUP BY r.name
    `).Scan(&roleCounts).Error; err != nil {
        return nil, err
    }
    for _, rc := range roleCounts {
        switch rc.Role {
        case "Admin":
            summary.TotalUsersAdmin = rc.Count
        case "Farmer":
            summary.TotalUsersFarmer = rc.Count
        }
        summary.TotalUsers += rc.Count
    }

    summary.TotalHarvestYield = 0

    return summary, nil
}