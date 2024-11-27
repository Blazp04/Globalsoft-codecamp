package infrastructure

import (
	"blazperic/lekcija10/port"
	"database/sql"
	"fmt"
	"time"
)

type SqliteAdapter struct {
	dbClient *sql.DB
}

func NewPersistenceAdapter(dbClient *sql.DB) *SqliteAdapter {
	return &SqliteAdapter{
		dbClient: dbClient,
	}
}

func (a *SqliteAdapter) GetTask(id int) (*port.TaskDTO, error) {
	findTaskSqlStatement := `
  SELECT title, description, deadline, completed, deleted FROM task WHERE id = ?;
`
	statement, err := a.dbClient.Prepare(findTaskSqlStatement)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare query: %v", err.Error())
	}
	defer statement.Close()

	var TaskDTO port.TaskDTO
	err = statement.
		QueryRow(id).
		Scan(&TaskDTO.Title, &TaskDTO.Description, &TaskDTO.Deadline, &TaskDTO.Completed, &TaskDTO.Deleted)
	if err != nil {
		return nil, fmt.Errorf("unable to set ID into statement: %v", err.Error())
	}

	return &TaskDTO, nil
}

func (a *SqliteAdapter) NewTask(title, description string, deadline time.Time, completed bool) error {
	createTaskSql := `INSERT INTO task(title, description, deadline, completed) values(?, ?, ?, ?)`

	stmt, err := a.dbClient.Prepare(createTaskSql)
	if err != nil {
		return fmt.Errorf("unable to prepare insert statement: %v", err.Error())
	}

	_, err = stmt.Exec(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("unable to execute insert statement: %v", err.Error())
	}
	return nil
}
