package repo

import (
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DimWebDev/task-manager-tool/internal/model"
)


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

    // Create a DueDate to use in your test
    dueDate := time.Now()

    // Use sqlmock.AnyArg() or a matcher that can match a time.Time for DueDate
    mock.ExpectExec("INSERT INTO tasks").
        WithArgs("Test Task", "This is a test task", sqlmock.AnyArg(), "Medium", "Pending").
        WillReturnResult(sqlmock.NewResult(1, 1))

    // Make sure to take the address of dueDate to get a *time.Time for DueDate
    task := model.Task{
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     &dueDate, // Now you're passing a *time.Time, assuming that's what your struct expects
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

    // Use a pointer to fixedTime in the mock response
    mock.ExpectQuery("SELECT id, title, description, duedate, priority, status FROM tasks WHERE id = \\$1").
        WithArgs(1).
        WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "duedate", "priority", "status"}).
            AddRow(1, "Test Task", "This is a test task", &fixedTime, "Medium", "Pending"))

    task, err := repo.GetByID(1)
    if err != nil {
        t.Errorf("error was not expected while getting task by ID: %s", err)
    }

    // Use a pointer to fixedTime in the expected result
    expected := model.Task{
        ID:          1,
        Title:       "Test Task",
        Description: "This is a test task",
        DueDate:     &fixedTime,
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

func TestGetAll(t *testing.T) {
    db, mock := NewMock()
    repo := NewTaskRepo(db)
    defer db.Close()

    // Creating a fixed time for consistency in tests
    fixedTime := time.Date(2024, 1, 10, 20, 50, 30, 0, time.UTC)

    // Mocking database response to return multiple rows of tasks
    rows := sqlmock.NewRows([]string{"id", "title", "description", "duedate", "priority", "status"}).
        AddRow(1, "Test Task 1", "This is the first test task", fixedTime, "High", "Pending").
        AddRow(2, "Test Task 2", "This is the second test task", fixedTime, "Medium", "Completed")

    mock.ExpectQuery("SELECT id, title, description, duedate, priority, status FROM tasks").
        WillReturnRows(rows)

    // Calling GetAll
    tasks, err := repo.GetAll()
    if err != nil {
        t.Errorf("error was not expected while getting all tasks: %s", err)
    }

    // Create pointers to fixedTime for the expected result
    fixedTimePtr1 := fixedTime
    fixedTimePtr2 := fixedTime

    // Define what we expect, using pointers for dates
    expected := []model.Task{
        {
            ID:          1,
            Title:       "Test Task 1",
            Description: "This is the first test task",
            DueDate:     &fixedTimePtr1,
            Priority:    "High",
            Status:      "Pending",
        },
        {
            ID:          2,
            Title:       "Test Task 2",
            Description: "This is the second test task",
            DueDate:     &fixedTimePtr2,
            Priority:    "Medium",
            Status:      "Completed",
        },
    }

    // Comparing the actual result with the expected result
    if !reflect.DeepEqual(tasks, expected) {
        t.Errorf("expected tasks %v, got %v", expected, tasks)
    }

    // Ensure all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}


func TestUpdate(t *testing.T) {
    db, mock := NewMock()
    repo := NewTaskRepo(db)
    defer db.Close()

    // Creating a fixed time for consistency in tests
    fixedTime := time.Date(2024, 1, 10, 20, 50, 30, 0, time.UTC)

    // Mocking the database to expect an UPDATE query with specific arguments
    // As we're passing fixedTime as a value, it is important to note that sqlmock will
    // match this based on the value passed, if your method sends it as a pointer,
    // you will need to match using sqlmock.AnyArg() instead.
    mock.ExpectExec("UPDATE tasks SET title = \\$1, description = \\$2, duedate = \\$3, priority = \\$4, status = \\$5 WHERE id = \\$6").
        WithArgs("Updated Test Task", "This is an updated test task", fixedTime, "High", "Completed", 1).
        WillReturnResult(sqlmock.NewResult(1, 1))

    // Creating a task struct with updated values
    // DueDate is a pointer to fixedTime
    updatedTask := model.Task{
        ID:          1,
        Title:       "Updated Test Task",
        Description: "This is an updated test task",
        DueDate:     &fixedTime,
        Priority:    "High",
        Status:      "Completed",
    }

    // Calling Update
    if err := repo.Update(updatedTask); err != nil {
        t.Errorf("error was not expected while updating task: %s", err)
    }

    // Ensure all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}

func TestDelete(t *testing.T) {
    db, mock := NewMock()
    repo := NewTaskRepo(db)
    defer db.Close()

    // Mocking the database to expect a DELETE query with a specific task ID
    mock.ExpectExec("DELETE FROM tasks WHERE id = \\$1").
        WithArgs(1).
        WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected

    // Calling Delete
    if err := repo.Delete(1); err != nil {
        t.Errorf("error was not expected while deleting task: %s", err)
    }

    // Ensure all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}



func TestMain(m *testing.M) {
	// Call flag.Parse() here if TestMain uses flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	m.Run()
}
