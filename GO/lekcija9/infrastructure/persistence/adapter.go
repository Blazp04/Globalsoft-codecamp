package persistence

import (
	"blazperic/lekcija9/core/port"
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
	//TODO selecting deadline causes error, check it out
	findTaskSqlStatement := `
  SELECT title, description, completed, deleted FROM task WHERE id = ?;
`
	statement, err := a.dbClient.Prepare(findTaskSqlStatement)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare query: %v", err.Error())
	}
	defer statement.Close()

	var TaskDTO port.TaskDTO
	err = statement.QueryRow(id).Scan(&TaskDTO.Title, &TaskDTO.Description, &TaskDTO.Completed, &TaskDTO.Deleted)
	if err != nil {
		return nil, fmt.Errorf("unable to set ID into statement: %v", err.Error())
	}

	return &TaskDTO, nil
}

func (a *SqliteAdapter) NewTask(title, description string, deadline time.Time, completed, deleted bool) error {
	insertTaskSqlStatement := `
  INSERT INTO task (title, description, deadline, completed, deleted)
  VALUES (?, ?, ?, ?, ?);
  `
	//Ianče radimo prepare statement jer se tako štitimo od sql injectiona
	_, err := a.dbClient.Exec(insertTaskSqlStatement, title, description, deadline, completed, deleted)
	if err != nil {
		return fmt.Errorf("unable to insert new task: %v", err.Error())
	}
	return nil
}
