/*
The package for CRUD-operations for Task Tracker App
*/
package crud

func (task *Task) Add(descriptrion []byte) error {}

func (task *Task) Update(new_descriptrion []byte) error {}

func (task *Task) Delete(id int) error {}

func (task *Task) Mark(id int, status []byte) error {}

func List(tasks *[]Task, status Status) error {}
