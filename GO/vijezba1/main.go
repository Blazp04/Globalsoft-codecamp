package main

import (
	"blazperic/santa/adapter"
	"blazperic/santa/database"
	"blazperic/santa/service"
	"fmt"
)

func main() {
	fmt.Println("Radi!!!")
	db, err := database.NewSqliteDatabase()

	if err != nil {
		panic(err)

	}
	err = db.Migrate()
	if err != nil {
		panic(err)

	}

	persistance := adapter.CreateDatabaseAdapter(db.GetDatabase())
	santaService := service.NewSantaService(persistance)
	// err = santaService.AddSanta("Antonija", false)
	// if err != nil {
	// 	panic(err)
	// }
	// err = santaService.AddSanta("Blaž", false)
	// if err != nil {
	// 	panic(err)
	// }
	santas, err := santaService.GetAllSantas()
	if err != nil {
		panic(err)
	}
	for _, a := range santas {
		fmt.Printf("Ime: %s, Izabrano: %t\n", a.Ime, a.IsChosen)
	}

	restController := adapter.NewRestController(santaService)
	restController.Run()

}
