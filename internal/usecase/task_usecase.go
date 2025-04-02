package usecase

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
)

type TaskUsecase struct {
	repo *repository.TaskRepository
}

func NewTaskUsecase(repo *repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo}
}

// Lấy danh sách công việc theo mùa vụ
func (u *TaskUsecase) GetTasksBySeason(seasonID uint) ([]models.Task, error) {
	return u.repo.GetTasksBySeason(seasonID)
}

// Tạo công việc mới
func (u *TaskUsecase) CreateTask(task *models.Task) error {
	return u.repo.CreateTask(task)
}

// Cập nhật trạng thái công việc
func (u *TaskUsecase) UpdateTaskStatus(taskID uint, status string) error {
	return u.repo.UpdateTaskStatus(taskID, status)
}
