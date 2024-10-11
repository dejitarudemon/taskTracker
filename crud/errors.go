/*
The package for CRUD-operations for Task Tracker App
*/
package crud

import "fmt"

type TaskError struct {
	id  int
	msg string
}

func (err *TaskError) Error() string {
	if err.id > 1 {
		return fmt.Sprintf("Msg: "+string(err.msg)+" with id %v", err.id)
	}
	return "Msg: " + string(err.msg)
}
