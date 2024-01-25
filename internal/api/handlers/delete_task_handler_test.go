package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// Assuming MockTaskRepository is already defined and implements repo.TaskRepository

func TestDeleteTaskHandler_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	taskID := 1
	mockRepo.On("Delete", taskID).Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(taskID)})
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code, "handler should return 204 No Content on success")

	mockRepo.AssertExpectations(t)
}

func TestDeleteTaskHandler_InvalidID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	req, _ := http.NewRequest("DELETE", "/tasks/invalidID", nil)
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "handler should return 400 Bad Request for invalid ID")

	mockRepo.AssertNotCalled(t, "Delete")
}

func TestDeleteTaskHandler_NotFound(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	taskID := 1
	mockRepo.On("Delete", taskID).Return(sql.ErrNoRows) // Simulate not found error

	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(taskID)})
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "handler should return 500 Internal Server Error if task not found")

	mockRepo.AssertExpectations(t)
}

