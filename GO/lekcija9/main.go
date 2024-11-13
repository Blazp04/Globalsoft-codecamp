package main

import (
	"fmt"
	"todo-cc/infrastructure/persistence"
	"todo-cc/infrastructure/rest"
	"todo-cc/persistence/sqlite"
)

func main() {
	restController := rest.NewRestController()
	db, err := sqlite.NewSqliteDatabase()
	if err != nil {
		panic(fmt.Sprintf("error while initializing database: %s", err.Error()))
	}
	err = db.MigrateDB()
	if err != nil {
		panic(err.Error())
	}
	//TODO remove this, test purpose only.
	taskPersistence := persistence.NewPersistenceAdapter(db.GetDb())
	_, _ = taskPersistence.GetTask(1)
	restController.Run()
}
