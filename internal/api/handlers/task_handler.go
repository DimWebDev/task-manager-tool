package handlers

import (
	"github.com/DimWebDev/task-manager-tool/internal/repo"
)

// TaskHandler holds the methods to handle task-related requests
type TaskHandler struct {
	Repo *repo.TaskRepo
}

// NewTaskHandler creates a new TaskHandler with the given repository
func NewTaskHandler(r *repo.TaskRepo) *TaskHandler {
	return &TaskHandler{Repo: r}
}
