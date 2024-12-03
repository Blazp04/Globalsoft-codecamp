package main

import (
	"blazperic/vjezba2/adapter"
	database "blazperic/vjezba2/databse"
	"blazperic/vjezba2/service"
	"fmt"
)

func main() {
	db, err := database.NewSqliteDatabase()
	if err != nil {
		panic(err)
	}
	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	persistanceAdapter := adapter.NewPersistanceAdapter(db.GetDatabase())
	noteService := service.NewNoteService(persistanceAdapter)
	restAdapter := adapter.NewRestAdapter(noteService)

	err = restAdapter.Run()
	if err != nil {
		panic(err)
	}

	_ = persistanceAdapter.NewNote("Naslov", "Opis")
	note, err := noteService.GetNote(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ime: %s, Opis: %s\n", note.Title, note.Description)

}
