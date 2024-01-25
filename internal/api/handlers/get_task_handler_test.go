package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DimWebDev/task-manager-tool/internal/model"
	"github.com/DimWebDev/task-manager-tool/internal/repo"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)



var _ repo.TaskRepository = &MockTaskRepository{}

func TestGetTaskByID(t *testing.T) {
    repoMock := new(MockTaskRepository)
    handler := NewTaskHandler(repoMock)

    router := mux.NewRouter()
    router.HandleFunc("/task/{id:[0-9]+}", handler.GetTaskByID)

    t.Run("Valid Task ID", func(t *testing.T) {
        dueDate := time.Now().Round(0) 
        task := model.Task{
            ID:          1,
            Title:       "Test Task",
            Description: "This is a test task",
            DueDate:     &dueDate,
            Priority:    "High",
            Status:      "Pending",
        }

        repoMock.On("GetByID", 1).Return(task, nil)

        req := httptest.NewRequest("GET", "/task/1", nil)
        rr := httptest.NewRecorder()

        router.ServeHTTP(rr, req)
        
        assert.Equal(t, http.StatusOK, rr.Code)
        
        var returnedTask model.Task
        err := json.Unmarshal(rr.Body.Bytes(), &returnedTask)
        assert.NoError(t, err)
        
        assert.Equal(t, task.ID, returnedTask.ID)
        assert.Equal(t, task.Title, returnedTask.Title)
        assert.Equal(t, task.Description, returnedTask.Description)
        assert.Equal(t, task.Priority, returnedTask.Priority)
        assert.Equal(t, task.Status, returnedTask.Status)
        assert.True(t, task.DueDate.Equal(*returnedTask.DueDate), "Due dates don't match")

        repoMock.AssertCalled(t, "GetByID", 1)
    })

    t.Run("Invalid Task ID", func(t *testing.T) {
        req := httptest.NewRequest("GET", "/task/abc", nil)
        rr := httptest.NewRecorder()

        router.ServeHTTP(rr, req)
        
        // Expect a 404 because the route does not match
        assert.Equal(t, http.StatusNotFound, rr.Code)
    })

    t.Run("Task Not Found", func(t *testing.T) {
        repoMock.On("GetByID", 99).Return(model.Task{}, errors.New("not found"))
        
        req := httptest.NewRequest("GET", "/task/99", nil)
        rr := httptest.NewRecorder()

        router.ServeHTTP(rr, req)
        
        assert.Equal(t, http.StatusNotFound, rr.Code)
        
        repoMock.AssertCalled(t, "GetByID", 99)
    })
}