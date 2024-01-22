// internal/api/handlers/delete_task_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteTask is an HTTP handler for deleting a task.
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// If the ID is not an integer, return a bad request response.
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Call the Delete method on the repository.
	err = h.Repo.Delete(id)
	if err != nil {
		// If there is an error deleting the task (e.g., task not found),
		// return an internal server error response.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the task was successfully deleted, return a no content response.
	w.WriteHeader(http.StatusNoContent)
}
