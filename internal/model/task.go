package model

import "time"

// Task represents a task structure.
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    Priority    string    `json:"priority"`
    Status      string    `json:"status"`
}
