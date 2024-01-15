package model

import "time"

type Task struct {
    ID          int       `json:"id,omitempty"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     *time.Time `json:"dueDate,omitempty"`
    Priority    string    `json:"priority,omitempty"`
    Status      string    `json:"status,omitempty"`
}