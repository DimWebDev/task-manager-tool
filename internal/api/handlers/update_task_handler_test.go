package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/DimWebDev/task-manager-tool/internal/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTaskHandler_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	// Assume we're updating task with ID 1
	taskID := 1
	task := model.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockRepo.On("Update", task).Return(nil)

	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(taskJSON))
	// Need to create a new router that has the route we are testing registered
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	rr := httptest.NewRecorder()

	// Need to mock the variables in the mux router.
	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(taskID),
	})

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Unmarshal the response body to check if the task returned matches the updated task
	var updatedTask model.Task
	err := json.Unmarshal(rr.Body.Bytes(), &updatedTask)
	assert.NoError(t, err, "should not error when unmarshaling the response")
	assert.Equal(t, taskID, updatedTask.ID, "handler returned incorrect ID")
	assert.Equal(t, task.Title, updatedTask.Title, "handler returned incorrect Title")
	assert.Equal(t, task.Description, updatedTask.Description, "handler returned incorrect Description")

	mockRepo.AssertExpectations(t)
}

func TestUpdateTaskHandler_FailInvalidID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	// Invalid ID
	req, _ := http.NewRequest("PUT", "/tasks/invalidID", nil)
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "should return 400 for invalid ID")
	// The mockRepo.Update method should not be called, so we expect no interactions with the mock
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateTaskHandler_FailInvalidBody(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	// Assume we're testing task with ID 1
	taskID := 1

	// Invalid JSON body
	invalidJSON := "{this is: invalid JSON}"
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBufferString(invalidJSON))
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(taskID),
	})

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "should return 400 for invalid JSON body")
	// The mockRepo.Update method should not be called, so we expect no interactions with the mock
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateTaskHandler_FailNotFound(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	handler := NewTaskHandler(mockRepo)

	// Assume we're updating task with ID 1
	taskID := 1
	task := model.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockRepo.On("Update", task).Return(sql.ErrNoRows) // Simulate task not found

	taskJSON, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(taskJSON))
	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{
		"id": strconv.Itoa(taskID),
	})

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "should return 404 when task not found")

	mockRepo.AssertExpectations(t)
}
