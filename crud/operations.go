/*
The package for CRUD-operations for Task Tracker App
*/
package crud

import "os"

const FILEDATA = "./tasks.json"

func (task *Task) Add(descriptrion []byte) error {
	file, err := os.Open(FILEDATA)
	if err != nil {
		return &TaskError{id: 0, msg: "No data file"}
	}
}

func (task *Task) Update(new_descriptrion []byte) error {}

func (task *Task) Delete(id int) error {}

func (task *Task) Mark(id int, status []byte) error {}

func List(tasks *[]Task, status Status) error {}
