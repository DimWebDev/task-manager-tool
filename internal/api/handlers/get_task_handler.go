package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetTaskByID is the handler for retrieving a task by its ID
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extracting the task ID from the URL path
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Task ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID format", http.StatusBadRequest)
		return
	}

	// Retrieve the task from the repository
	task, err := h.Repo.GetByID(id)
	if err != nil {
		// Handle the case where the task is not found
		http.Error(w, "Task not found", http.StatusNotFound)
		return
		// Handle other potential errors
		// Assume that repo.GetByID will return a sql.ErrNoRows for not found tasks
		// and other errors for different unexpected situations.
	}

	// Respond with the task in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
	}
}
