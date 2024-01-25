package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DimWebDev/task-manager-tool/internal/model"
	"github.com/stretchr/testify/mock"
)



func TestCreateTaskHandler_ValidInput(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	testHandler := NewTaskHandler(mockRepo)

    dueDate := time.Now().Add(24 * time.Hour) 
    task := model.Task{
        Title:       "Test Task",
        Description: "Test Description",
        DueDate:     &dueDate,
        Priority:    "High",
        Status:      "New",
    }

    mockRepo.On("Create", mock.AnythingOfType("model.Task")).Return(nil)

    taskJSON, _ := json.Marshal(task)
    req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
    rr := httptest.NewRecorder()

    testHandler.CreateTaskHandler(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    // Unmarshal the response body into a Task object for comparison
    var returnedTask model.Task
    if err := json.Unmarshal(rr.Body.Bytes(), &returnedTask); err != nil {
        t.Fatalf("could not unmarshal response body: %v", err)
    }

    // Compare all fields.
    if returnedTask.Title != task.Title {
        t.Errorf("handler returned unexpected title: got %v want %v", returnedTask.Title, task.Title)
    }
    if returnedTask.Description != task.Description {
        t.Errorf("handler returned unexpected description: got %v want %v", returnedTask.Description, task.Description)
    }
    if returnedTask.Priority != task.Priority {
        t.Errorf("handler returned unexpected priority: got %v want %v", returnedTask.Priority, task.Priority)
    }
    if returnedTask.Status != task.Status {
        t.Errorf("handler returned unexpected status: got %v want %v", returnedTask.Status, task.Status)
    }
    // Compare the DueDate separately as it's a pointer to time.Time.
    if returnedTask.DueDate == nil || !returnedTask.DueDate.Equal(*task.DueDate) {
        t.Errorf("handler returned unexpected dueDate: got %v want %v", returnedTask.DueDate, task.DueDate)
    }

    mockRepo.AssertExpectations(t)
}

func TestCreateTaskHandler_InvalidJSON(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	testHandler := NewTaskHandler(mockRepo)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBufferString("{invalid JSON"))
	rr := httptest.NewRecorder()

	testHandler.CreateTaskHandler(rr, req)

	// Check that the status code is 400 BadRequest
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Ensure no expectations were met as the handler should have responded before calling repo.Create
	mockRepo.AssertNotCalled(t, "Create")
}
