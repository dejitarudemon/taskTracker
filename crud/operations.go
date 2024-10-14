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

func Init() error {
	_, err := os.Stat(FILEDATA)

	if err == nil {
		return nil
	}

	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(FILEDATA)
		return err
	}

	return err
}

func Add(descriptrion string) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	var last_id int
	if len(tasks) > 0 {
		last_id = tasks[len(tasks)-1].id
	} else {
		last_id = 1
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

func Update(task_id int, new_descriptrion string) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	task := find(&tasks, task_id)
	if task == nil {
		return errors.New("NO TASK WITH THE ID")
	}

	task.description = new_descriptrion
	task.updatedAt = time.Now()
	return dump(&tasks, FILEDATA)
}

func Delete(task_id int) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	task := find(&tasks, task_id)
	if task == nil {
		return errors.New("NO TASK WITH THE ID")
	}

	new_tasks := make([]Task, 0, len(tasks)-1)
	for i := range tasks {
		if tasks[i].id != task.id {
			new_tasks = append(new_tasks, tasks[i])
		}
	}

	return dump(&new_tasks, FILEDATA)
}

func Mark(task_id int, status Status) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	task := find(&tasks, task_id)
	if task == nil {
		return errors.New("NO TASK WITH THE ID")
	}

	task.status = status
	return dump(&tasks, FILEDATA)
}

func List(status *Status) ([]Task, error) {
	tasks, err := load(FILEDATA)
	if err != nil {
		return []Task{}, err
	}

	if status != nil {
		return filter(&tasks, status), nil
	} else {
		return tasks, nil
	}
}

func load(filepath string) ([]Task, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []Task{}, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return []Task{}, err
	}

	if !json.Valid(content) {
		return []Task{}, errors.New("NOT JSON-VALIDE")
	}

	var tasks []Task

	if err := json.Unmarshal(content, &tasks); err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func dump(tasks *[]Task, filepath string) error {
	file, err := os.Open(filepath)
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

func find(tasks *[]Task, task_id int) *Task {
	for _, task := range *tasks {
		if task.id == task_id {
			return &task
		}
	}
	return nil
}

func filter(tasks *[]Task, status *Status) []Task {
	result := make([]Task, 0, 1)

	for _, task := range *tasks {
		if task.status == *status {
			result = append(result, task)
		}
	}

	return result
}
