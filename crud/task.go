/*
The package for CRUD-operations for Task Tracker App
*/
package crud

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Tasks []Task

const (
	Done       = "done"
	InProgress = "in-progress"
	ToDo       = "todo"
)
