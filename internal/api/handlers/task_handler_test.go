package handlers

import (
	"testing"

	"github.com/DimWebDev/task-manager-tool/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
    mock.Mock
}

func (m *MockTaskRepository) GetByID(id int) (model.Task, error) {
    args := m.Called(id)
    return args.Get(0).(model.Task), args.Error(1)
}

func (m *MockTaskRepository) Create(task model.Task) error {
    args := m.Called(task)
    return args.Error(0)
}

func (m *MockTaskRepository) GetAll() ([]model.Task, error) {
    args := m.Called()
    return args.Get(0).([]model.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task model.Task) error {
    args := m.Called(task)
    return args.Error(0)
}

func (m *MockTaskRepository) Delete(id int) error {
    args := m.Called(id)
    return args.Error(0)
}

func TestNewTaskHandler(t *testing.T) {
    // Instantiate the mock repository
    mockRepo := new(MockTaskRepository)
    
    // Call NewTaskHandler using the mock repo instance
    taskHandler := NewTaskHandler(mockRepo)
    
    // Check that the returned TaskHandler is not nil
    if taskHandler == nil {
        t.Error("NewTaskHandler returned nil")
    }
    
    // Check that the Repo field is set correctly
    if taskHandler.Repo != mockRepo {
        t.Errorf("Expected TaskHandler.Repo to be %v, got %v", mockRepo, taskHandler.Repo)
    }
}