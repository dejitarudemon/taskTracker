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
		file, err := os.Create(FILEDATA)
		file.Close()
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
		last_id = tasks[len(tasks)-1].Id
	} else {
		last_id = 0
	}

	task := Task{
		Id:          last_id + 1,
		Description: descriptrion,
		Status:      ToDo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	tasks = append(tasks, task)

	return dump(&tasks, FILEDATA)
}

func Update(task_id int, new_descriptrion string) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Id == task_id {
			tasks[i].Description = new_descriptrion
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return dump(&tasks, FILEDATA)
		}
	}
	return errors.New("NO TASK WITH THE ID")
}

func Delete(task_id int) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return errors.New("NO TASK WITH THE ID")
	}

	new_tasks := make(Tasks, 0, len(tasks)-1)
	for i := range tasks {
		if tasks[i].Id != task_id {
			new_tasks = append(new_tasks, tasks[i])
		}
	}

	if len(new_tasks) == len(tasks) {
		return errors.New("NO TASK WITH THE ID")
	}
	return dump(&new_tasks, FILEDATA)
}

func Mark(task_id int, status string) error {
	tasks, err := load(FILEDATA)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].Id == task_id {
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			tasks[i].Status = status
			return dump(&tasks, FILEDATA)
		}
	}
	return errors.New("NO TASK WITH THE ID")
}

func List(status *string) (Tasks, error) {
	tasks, err := load(FILEDATA)
	if err != nil {
		return Tasks{}, err
	}

	if status != nil {
		return filter(&tasks, status), nil
	} else {
		return tasks, nil
	}
}

func load(filepath string) (Tasks, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Tasks{}, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return Tasks{}, err
	}

	if len(content) == 0 {
		return Tasks{}, nil
	}

	var tasks Tasks

	if err := json.Unmarshal(content, &tasks); err != nil {
		return Tasks{}, err
	}

	return tasks, nil
}

func dump(tasks *Tasks, filepath string) error {
	contentJsoned, err := json.MarshalIndent(*tasks, "\t", "\n")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath, contentJsoned, 0664); err != nil {
		return err
	}
	return nil
}

func filter(tasks *Tasks, status *string) Tasks {
	result := make(Tasks, 0, 1)

	for _, task := range *tasks {
		if task.Status == *status {
			result = append(result, task)
		}
	}

	return result
}
