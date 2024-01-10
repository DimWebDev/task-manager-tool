package model

import "time"

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"duedate"`
    Priority    string    `json:"priority"`
    Status      string    `json:"status"`
}
