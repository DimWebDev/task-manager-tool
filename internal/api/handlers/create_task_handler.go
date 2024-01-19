// create_task_handler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DimWebDev/task-manager-tool/internal/model"
)

// CreateTaskHandler handles the creation of a new task
func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask model.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid task format", http.StatusBadRequest)
		return
	}

	// Validate the task as needed
	// For example: check if the title is not empty
	if newTask.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Call the repository function to insert the new task
	err = h.Repo.Create(newTask)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// If creation is successful, return the new task with StatusCreated
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
