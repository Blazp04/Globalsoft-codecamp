package sqllite

import (
	"blazperic/lekcija9/config"
	"blazperic/lekcija9/persistance"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Db struct {
	database *sql.DB
}

func NewSqlDatabase() (*Db, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.DATABASE_NAME)
	if err != nil {
		// _ = fmt.Errorf("error opening database: %v", err)
		return nil, &persistance.DbConnectionError{Message: fmt.Sprintf("error opening database: %v", err.Error())}
	}
	dbInstance := &Db{database: db}
	return dbInstance, nil
}

func (sl *Db) GetDB() *sql.DB {
	return sl.database
}

func (sl *Db) Migrate() error {
	initialSqlStatement := "CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, title TEXT NOT NULL, description TEXT, deadline DATE, isCompleted BOOLEAN DEFAULT FALSE, isDeleted BOOLEAN DEFAULT FALSE);"

	_, err := sl.GetDB().Exec(initialSqlStatement)
	if err != nil {
		return &persistance.DbStatementError{Message: fmt.Sprintf("error creating table: %v", err.Error())}
	}
	return nil
}
