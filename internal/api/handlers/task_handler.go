package handlers

import (
	"github.com/DimWebDev/task-manager-tool/internal/repo"
)

// TaskHandler holds the methods to handle task-related requests. Each of these method is defined inside the specific handler files
type TaskHandler struct {
    Repo repo.TaskRepository
}

// NewTaskHandler creates a new TaskHandler with the given repository
func NewTaskHandler(r repo.TaskRepository) *TaskHandler {
    return &TaskHandler{Repo: r}
}
