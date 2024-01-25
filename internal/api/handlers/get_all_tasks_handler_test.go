package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DimWebDev/task-manager-tool/internal/model"
	"github.com/stretchr/testify/assert"
)



func TestGetAllTasks(t *testing.T) {
    repoMock := new(MockTaskRepository)
    handler := NewTaskHandler(repoMock)

    tasks := []model.Task{
        {
            ID:          1,
            Title:       "Task 1",
            Description: "Description 1",
            DueDate:     &time.Time{},
            Priority:    "High",
            Status:      "New",
        },
        {
            ID:          2,
            Title:       "Task 2",
            Description: "Description 2",
            DueDate:     &time.Time{},
            Priority:    "Medium",
            Status:      "In Progress",
        },
    }

    repoMock.On("GetAll").Return(tasks, nil)

    req := httptest.NewRequest("GET", "/tasks", nil)
    rr := httptest.NewRecorder()

    handler.GetAllTasks(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)
    assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

    var returnedTasks []model.Task
    err := json.Unmarshal(rr.Body.Bytes(), &returnedTasks)
    assert.NoError(t, err)
    assert.Equal(t, tasks, returnedTasks)

    repoMock.AssertExpectations(t)
}
