/*
The package for CRUD-operations for Task Tracker App
*/
package crud

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

const FILEDATA = "./tasks.json"

func Add(descriptrion []byte) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	var last_id int
	if len(tasks) > 0 {
		last_id = tasks[len(tasks)-1].id
	} else {
		last_id = 0
	}

	task := Task{
		id:          last_id + 1,
		description: descriptrion,
		status:      Status(ToDo),
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	return dump(&tasks, FILEDATA)
}

func Update(task_id int, new_descriptrion []byte) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.id == task_id {
			task.description = new_descriptrion
			return dump(&tasks, FILEDATA)
		}
	}

	return errors.New("There is no task with the id")
}

func (task *Task) Delete(id int) error {}

func (task *Task) Mark(id int, status []byte) error {}

func List(tasks *[]Task, status Status) error {}

func load(filepath string) ([]Task, error) {
	file, err := os.Open(FILEDATA)
	if err != nil {
		return []Task{}, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return []Task{}, err
	}

	if !json.Valid(content) {
		return []Task{}, errors.New("Not Valid Structure")
	}

	var tasks []Task

	if err := json.Unmarshal(content, &tasks); err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func dump(tasks *[]Task, filepath string) error {
	file, err := os.Open(FILEDATA)
	if err != nil {
		return err
	}
	defer file.Close()

	contentJsoned, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if _, err := file.Write(contentJsoned); err != nil {
		return err
	}
	return nil
}
