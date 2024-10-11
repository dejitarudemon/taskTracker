/*
The package for CRUD-operations for Task Tracker App
*/
package crud

func Add(descriptrion []byte) error {}

func Update(id int, new_descriptrion []byte) error {}

func Delete(id int) error {}

func Mark(id int, status []byte) error {}

func List(id int, status Status) error {}
