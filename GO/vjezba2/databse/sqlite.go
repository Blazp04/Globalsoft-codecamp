package database

import (
	"blazperic/vjezba2/config"
	"blazperic/vjezba2/shared"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Db struct {
	database *sql.DB
}

func NewSqliteDatabase() (*Db, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.DATABASE_NAME)
	if err != nil {
		return nil, &shared.DatabaseOpeningError{
			Message: fmt.Sprintf("Dogodila se greška pri otvaranju baze podataka: %s", err.Error()),
		}
	}
	return &Db{database: db}, nil
}

func (s *Db) GetDatabase() *sql.DB {
	return s.database
}

func (s *Db) Migrate() error {
	sqlStatement := "CREATE TABLE IF NOT EXISTS notes(id INTEGER PRIMARY KEY, title TEXT NOT NULL, content TEXT NOT NULL);"
	_, err := s.GetDatabase().Exec(sqlStatement)
	if err != nil {
		return &shared.ExecError{
			Message: fmt.Sprintf("error while migrating db: %v", err.Error())}
	}
	return nil
}
