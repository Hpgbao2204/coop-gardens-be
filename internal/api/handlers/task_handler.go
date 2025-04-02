package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	usecase *usecase.TaskUsecase
}

func NewTaskHandler(usecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase}
}

func (h *TaskHandler) GetTasksBySeason(c echo.Context) error {
	seasonID, err := strconv.Atoi(c.Param("season_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid season ID"})
	}

	tasks, err := h.usecase.GetTasksBySeason(uint(seasonID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := h.usecase.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTaskStatus(c echo.Context) error {
	taskID, _ := strconv.Atoi(c.Param("task_id"))
	var payload struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := h.usecase.UpdateTaskStatus(uint(taskID), payload.Status); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task status"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Task updated successfully"})
}
