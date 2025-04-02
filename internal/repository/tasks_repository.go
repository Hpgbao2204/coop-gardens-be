package repository

import (
	"coop-gardens-be/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTasksBySeason(seasonID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Preload("Crops").Where("season_id = ?", seasonID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) UpdateTaskStatus(taskID uint, status string) error {
	return r.DB.Model(&models.Task{}).Where("id = ?", taskID).Update("status", status).Error
}
