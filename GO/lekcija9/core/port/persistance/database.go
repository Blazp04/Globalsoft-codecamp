package persistance

import "database/sql"

type Database interface {
	GetDb() *sql.DB
}
