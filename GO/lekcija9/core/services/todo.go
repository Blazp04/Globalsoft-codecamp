package service

import "blazperic/lekcija9/core/port/persistance"

type Todo struct {
	persistence persistance.Port
}

func NewTodo(todoPersistence persistance.Port) Todo {
	return Todo{
		todoPersistence,
	}
}

func (t Todo) CreateNewTask() {

}
