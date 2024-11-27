package service

import (
	"blazperic/lekcija10/port"
	"time"
)

type Todo struct {
	persistence port.PersistencePort
}

func NewTodo(todoPersistence port.PersistencePort) Todo {
	return Todo{
		todoPersistence,
	}
}

func (t Todo) CreateNewTask(title, description string, deadline time.Time, completed bool) {
	err := t.persistence.NewTask(title, description, deadline, completed)
	if err != nil {
	}
}
