/*
The package for CRUD-operations for Task Tracker App
*/
package crud

import "time"

type Task struct {
	id          int
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type Status string

const (
	Done       = "done"
	InProgress = "in-progress"
	ToDo       = "todo"
)
