package main

import (
	"blazperic/lekcija9/infrastructure"
	"blazperic/lekcija9/persistance/sqllite"
	"fmt"
)

func main() {
	controller := infrastructure.NewRestController()

	db, err := sqllite.NewSqlDatabase()
	if err != nil {
		panic(fmt.Sprintf("error creating database: %v", err))
	}

	err = db.GetDB().Ping()
	if err != nil {
		panic(fmt.Sprintf("error pinging database: %v", err))
	}

	err = db.Migrate()
	if err != nil {
		panic(fmt.Sprintf("error migrating database: %v", err))
	}

	controller.Run()

	fmt.Println("Ovo sve radi!")
}
