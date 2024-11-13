package port

import "time"

type Port interface {
	GetTask(id int) (*TaskDTO, error)
	NewTask(title, description string, deadline time.Time, completed, deleted bool) error
}

type TaskDTO struct {
	Title       string
	Description string
	Deadline    time.Time
	Completed   bool
	Deleted     bool
}
