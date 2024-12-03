package adapter

import (
	"blazperic/vjezba2/port"
	"blazperic/vjezba2/shared"
	"database/sql"
	"fmt"
)

type PersistanceAdapter struct {
	dbClient *sql.DB
}

func NewPersistanceAdapter(dbClient *sql.DB) *PersistanceAdapter {
	return &PersistanceAdapter{
		dbClient: dbClient,
	}
}

func (p *PersistanceAdapter) NewNote(title, description string) error {
	sqlStatement := "INSERT INTO notes(title, content) VALUES(?, ?);"
	statement, err := p.dbClient.Prepare(sqlStatement)
	if err != nil {
		return &shared.ExecError{
			Message: fmt.Sprintf("Dogodila se greška prilikom stvaranja statementa za novi bilješku: %s", err.Error()),
		}
	}
	defer statement.Close()
	_, err = statement.Exec(title, description)
	if err != nil {
		return &shared.ExecError{
			Message: fmt.Sprintf("Dogodila se greška prilikom spremanja bilješke: %s", err.Error()),
		}
	}
	return nil
}

func (p *PersistanceAdapter) GetNote(id int) (*port.NoteDTO, error) {
	sqlStatement := "SELECT id, title, content FROM notes WHERE id = ?;"
	statement, err := p.dbClient.Prepare(sqlStatement)
	if err != nil {
		return nil, &shared.QeuryError{
			Message: fmt.Sprintf("Dogodila se greška prilikom stvaranja statementa za preuzimanje bilješke: %s", err.Error()),
		}
	}
	defer statement.Close()

	var note port.NoteDTO
	err = statement.QueryRow(id).Scan(&note.Id, &note.Title, &note.Description)
	if err != nil {
		return nil, &shared.QeuryError{
			Message: fmt.Sprintf("Dogodila se greška prilikom preuzimanja bilješke: %s", err.Error()),
		}
	}
	return &note, nil
}

func (p *PersistanceAdapter) GetAllNotes() ([]*port.NoteDTO, error) {
	sqlStatement := "SELECT id, title, content FROM notes;"

	rows, err := p.dbClient.Query(sqlStatement)
	if err != nil {
		return nil, &shared.QeuryError{
			Message: fmt.Sprintf("Dogodila se greška prilikom preuzimanja svih bilješki: %s", err.Error()),
		}
	}
	defer rows.Close()

	var notes []*port.NoteDTO

	for rows.Next() {
		var note port.NoteDTO
		err := rows.Scan(&note.Id, &note.Title, &note.Description)
		if err != nil {
			return nil, &shared.QeuryError{
				Message: fmt.Sprintf("Dogodila se greška prilikom preuzimanja svih bilješki: %s", err.Error()),
			}
		}
		notes = append(notes, &note)
	}
	return notes, nil
}
