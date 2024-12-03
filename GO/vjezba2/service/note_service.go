package service

import "blazperic/vjezba2/port"

type NoteService struct {
	Persistance port.Persistance
}

func NewNoteService(p port.Persistance) *NoteService {
	return &NoteService{Persistance: p}
}

func (s *NoteService) GetNote(id int) (*port.NoteDTO, error) {
	return s.Persistance.GetNote(id)
}

func (s *NoteService) GetAllNotes() ([]*port.NoteDTO, error) {
	return s.Persistance.GetAllNotes()
}

func (s *NoteService) NewNote(title, description string) error {
	return s.Persistance.NewNote(title, description)
}
