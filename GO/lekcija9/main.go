package main

import (
	"blazperic/lekcija9/infrastructure/persistence"
	"blazperic/lekcija9/infrastructure/rest"
	"blazperic/lekcija9/persistence/sqlite"
	"fmt"
	"time"
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
	_ = taskPersistence.NewTask("test", "test", time.Now().Add(time.Hour*24), false, false)
	res, _ := taskPersistence.GetTask(1)
	fmt.Println(res)

	restController.Run()
}
