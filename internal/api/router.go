package api

import (
	"net/http"

	"github.com/DimWebDev/task-manager-tool/internal/api/handlers"
	"github.com/gorilla/mux"
)




func NewRouter(taskHandler *handlers.TaskHandler) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/tasks", taskHandler.CreateTaskHandler).Methods(http.MethodPost)


	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTaskByID).Methods(http.MethodGet)

	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods(http.MethodPut)

	router.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	router.HandleFunc("/tasks", taskHandler.GetAllTasks).Methods(http.MethodGet)



	return router
}


