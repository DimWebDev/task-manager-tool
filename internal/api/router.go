package api

import (
	"net/http"

	"github.com/DimWebDev/task-manager-tool/internal/api/handlers"
	"github.com/gorilla/mux"
)

// NewRouter initializes and returns a new router with all the necessary routes defined.


func NewRouter(taskHandler *handlers.TaskHandler) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/tasks", taskHandler.CreateTaskHandler).Methods(http.MethodPost)

	// Routes for operations on a task by ID
	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskByID).Methods(http.MethodGet)


	// router.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	// router.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTaskHandler).Methods(http.MethodDelete)

	// The router is now set up and ready to be used
	return router
}


