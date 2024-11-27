package port

import "time"

type PersistencePort interface {
	GetTask(id int) (*TaskDTO, error)
	NewTask(title, description string, deadline time.Time, completed bool) error
	GetAllTasks() ([]*TaskDTO, error)
	DeleteTask(id int) error
	CompleteTask(id int) error
}

type TaskDTO struct {
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
	Deleted     bool
}
