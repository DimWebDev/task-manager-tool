// internal/handlers/task_handler.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DimWebDev/task-manager-tool/internal/model"
	"github.com/gorilla/mux"
)

// UpdateTask updates an existing task.
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from the URL.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a Task struct.
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Set the task ID from the URL.
	task.ID = id

	// Call the Update method on the repo.
	if err := h.Repo.Update(task); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// If successful, encode and return the updated task.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}
