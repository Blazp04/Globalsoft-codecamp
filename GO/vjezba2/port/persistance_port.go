package port

type Persistance interface {
	GetNote(id int) (*NoteDTO, error)
	NewNote(title, description string) error
	GetAllNotes() ([]*NoteDTO, error)
}

type NoteDTO struct {
	Id          int
	Title       string
	Description string
}
