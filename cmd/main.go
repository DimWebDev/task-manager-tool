// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DimWebDev/task-manager-tool/internal/api"
	myhandlers "github.com/DimWebDev/task-manager-tool/internal/api/handlers"
	"github.com/DimWebDev/task-manager-tool/internal/repo"
	_ "github.com/lib/pq"
)

const (
    host   = "localhost"
    port   = 5432
    dbname = "task_manager"
)

func main() {
    // Retrieve the username and password from the environment variables
    user := os.Getenv("POSTGRES_USER")
  
    if user == ""  {
        log.Fatal("The POSTGRES_USER environment variable is not set.")
    }

    // Construct the connection string
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "dbname=%s sslmode=disable", host, port, user, dbname)

    // Open a connection to the database
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening connection: %s", err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %s", err)
    }
    fmt.Println("Successfully connected to database!")

    // Initialize the repository
    taskRepo := repo.NewTaskRepo(db)

    // Initialize the handler with the repository
    taskHandler := myhandlers.NewTaskHandler(taskRepo)

    // Set up the router with the task handler
    router := api.NewRouter(taskHandler)

    // Start the HTTP server with the router
    httpAddress := ":8080"
    fmt.Printf("Starting server on %s\n", httpAddress)
    log.Fatal(http.ListenAndServe(httpAddress, router))
}