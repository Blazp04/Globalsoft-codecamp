package sqllite

import (
	"blazperic/lekcija9/persistance"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	database *sql.DB
}

func NewSqlDatabase() (*Db, error) {
	db, err := sql.Open("sqlite3", "test.db")
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
