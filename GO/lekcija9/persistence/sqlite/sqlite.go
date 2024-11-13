package sqlite

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"todo-cc/config"
	"todo-cc/persistence"
)

type Db struct {
	database *sql.DB
}

func NewSqliteDatabase() (*Db, error) {
	db, err := sql.Open(config.DRIVER_NAME, config.DATABASE_NAME)
	if err != nil {
		return nil, &persistence.DbConnectionError{
			Message: fmt.Sprintf("error while conneting to database: %s", err.Error()),
		}
	}
	dbInstance := &Db{database: db}

	return dbInstance, nil
}

func (sl *Db) GetDb() *sql.DB {
	return sl.database
}

func (sl *Db) MigrateDB() error {
	initialSqlStatement := `CREATE TABLE IF NOT EXISTS task (
id INT PRIMARY KEY,
title TEXT NOT NULL,
description TEXT,
deadline DATE,
completed BOOLEAN DEFAULT FALSE,
deleted BOOLEAN DEFAULT FALSE
);`
	_, err := sl.GetDb().Exec(initialSqlStatement)
	if err != nil {
		return &persistence.ExecError{
			Message: fmt.Sprintf("error while migrating db: %v", err.Error()),
		}
	}

	return nil
}
