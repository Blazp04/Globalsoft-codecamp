package service

import (
	"blazperic/lekcija10/port"
	"fmt"
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

func (t Todo) CreateNewTask(title, description string, deadline time.Time, completed bool) error {
	err := t.persistence.NewTask(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("error while saving task: %v+", err.Error())
	}

	return nil
}

func (t Todo) GetTask(id int) (*port.TaskDTO, error) {
	task, err := t.persistence.GetTask(id)
	if err != nil {
		return nil, err
	}
	return task, nil

}

func (t Todo) GetAllTasks() ([]*port.TaskDTO, error) {
	task, err := t.persistence.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t Todo) DeleteTask(id int) error {
	err := t.persistence.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (t Todo) CompleteTask(id int) error {
	err := t.persistence.CompleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
