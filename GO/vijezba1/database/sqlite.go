package database

import (
	"blazperic/santa/config"
	"blazperic/santa/shared"
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
		return nil, &shared.DbConnectionError{
			Message: fmt.Sprintf("error while conneting to database: %s", err.Error()),
		}
	}
	dbInstance := &Db{database: db}

	return dbInstance, nil
}

func (d *Db) GetDatabase() *sql.DB {
	return d.database
}

func (d *Db) Migrate() error {
	sqlStatement := "CREATE TABLE IF NOT EXISTS santas(id INTEGER PRIMARY KEY,ime TEXT NOT NULL, isChosen  BOOLEAN DEFAULT FALSE);"
	_, err := d.GetDatabase().Exec(sqlStatement)

	if err != nil {
		return err
	}
	return nil

}
