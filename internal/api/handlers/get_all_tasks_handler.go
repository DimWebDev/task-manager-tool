// internal/api/handlers/get_all_tasks_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
)

// GetAllTasks handles the HTTP request for retrieving all tasks.
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// Invoke the GetAll method to retrieve tasks
	tasks, err := h.Repo.GetAll()
	if err != nil {
		// If an error occurs, send an internal server error response
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type as application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the HTTP status code
	w.WriteHeader(http.StatusOK)
	// Encode and send the tasks as a JSON response
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
	}
}
