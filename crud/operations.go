/*
The package for CRUD-operations for Task Tracker App
*/
package crud

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const FILEDATA = "./tasks.json"

func (task *Task) Add(descriptrion []byte) error {
	file, err := os.Open(FILEDATA)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if !json.Valid(content) {
		return errors.New("Not Valid Structure")
	}

	var tasks []Task

	if err := json.Unmarshal(content, &tasks); err != nil {
		return err
	}

	var last_id int
	if len(tasks) > 0 {
		last_id = tasks[len(tasks)-1].id
	} else {
		last_id = 0
	}

	task.id = last_id + 1
	tasks = append(tasks, *task)

	tasksJsoned, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if _, err := file.Write(tasksJsoned); err != nil {
		return err
	}
	return nil
}

func (task *Task) Update(new_descriptrion []byte) error {}

func (task *Task) Delete(id int) error {}

func (task *Task) Mark(id int, status []byte) error {}

func List(tasks *[]Task, status Status) error {}
