package main

import (
	"blazperic/lekcija9/persistance/sqllite"
	"fmt"
)

func main() {
	db, err := sqllite.NewSqlDatabase()
	if err != nil {
		panic(fmt.Sprintf("error creating database: %v", err))
	}

	err = db.GetDB().Ping()
	if err != nil {
		panic(fmt.Sprintf("error pinging database: %v", err))
	}
	fmt.Println("Hello, World!")
}
