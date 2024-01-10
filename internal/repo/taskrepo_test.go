package repo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	// pq is the library that allows Go to interact with PostgreSQL
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DimWebDev/task-manager-tool/internal/model"
	_ "github.com/lib/pq"
)

const (
    host   = "localhost"
    port   = 5432
    dbname = "task_manager_test" // Use the test database
)

var db *sql.DB

func TestMain(m *testing.M) {
    // Retrieve the username from the environment variable
    user := os.Getenv("POSTGRES_USER")
    if user == "" {
        log.Fatal("The POSTGRES_USER environment variable is not set.")
    }

    // Construct the connection string for the test database
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "dbname=%s sslmode=disable", host, port, user, dbname)

    // Open a connection to the test database
    var err error
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening connection to test database: %s", err)
    }
    defer db.Close()

    // Test the connection to the test database
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging test database: %s", err)
    }

    // Run the tests
    code := m.Run()

    // Exit with the test exit code
    os.Exit(code)
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    if err != nil {
        log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    return db, mock
}


func TestCreateTask(t *testing.T) {
    db, mock := NewMock()
    repo := NewTaskRepo(db)
    defer db.Close()

    // Using sqlmock.AnyArg() for the time argument
    mock.ExpectExec("INSERT INTO tasks").
        WithArgs("Test Task", "This is a test task", sqlmock.AnyArg(), "Medium", "Pending").
        WillReturnResult(sqlmock.NewResult(1, 1))

    task := model.Task{
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     time.Now(),
        Priority:    "Medium",
        Status:      "Pending",
    }

    if err := repo.Create(task); err != nil {
        t.Errorf("error was not expected while creating task: %s", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestGetByID(t *testing.T) {
    db, mock := NewMock()
    repo := NewTaskRepo(db)
    defer db.Close()

    fixedTime := time.Date(2024, 1, 10, 20, 50, 30, 0, time.UTC) // Example fixed time

    mock.ExpectQuery("SELECT id, title, description, duedate, priority, status FROM tasks WHERE id = \\$1").
        WithArgs(1).
        WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "duedate", "priority", "status"}).
            AddRow(1, "Test Task", "This is a test task", fixedTime, "Medium", "Pending"))

    task, err := repo.GetByID(1)
    if err != nil {
        t.Errorf("error was not expected while getting task by ID: %s", err)
    }

    expected := model.Task{
        ID:          1,
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     fixedTime,
        Priority:    "Medium",
        Status:      "Pending",
    }

    if !reflect.DeepEqual(task, expected) {
        t.Errorf("expected task %v, got %v", expected, task)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}



