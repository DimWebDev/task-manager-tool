// internal/repo/taskrepo.go
// The taskrepo.go serves as the data access layer in your application architecture.
// It abstracts the CRUD operations related to tasks, directly interacting with the
// database without any business logic or presentation concerns.
package repo

import (
	"database/sql"

	"github.com/DimWebDev/task-manager-tool/internal/model"
)

// TaskRepo provides access to the task storage.
type TaskRepo struct {
	db *sql.DB
}

// NewTaskRepo creates a new TaskRepo.
func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

// Create inserts a new task into the database.
func (tr *TaskRepo) Create(task model.Task) error {
	_, err := tr.db.Exec("INSERT INTO tasks (title, description, duedate, priority, status) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Description, task.DueDate, task.Priority, task.Status)
	return err
}

// GetByID retrieves a task by its ID from the database.
func (tr *TaskRepo) GetByID(id int) (model.Task, error) {
	var task model.Task
	err := tr.db.QueryRow("SELECT id, title, description, due_date, priority, status FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

// GetAll retrieves all tasks from the database.
func (tr *TaskRepo) GetAll() ([]model.Task, error) {
	rows, err := tr.db.Query("SELECT id, title, description, due_date, priority, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Update modifies an existing task in the database.
func (tr *TaskRepo) Update(task model.Task) error {
	_, err := tr.db.Exec("UPDATE tasks SET title = $1, description = $2, due_date = $3, priority = $4, status = $5 WHERE id = $6",
		task.Title, task.Description, task.DueDate, task.Priority, task.Status, task.ID)
	return err
}

// Delete removes a task by its ID from the database.
func (tr *TaskRepo) Delete(id int) error {
	_, err := tr.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
