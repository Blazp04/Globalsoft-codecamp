package port

type Database interface {
	AddSanta(ime string, isChosen bool) error
	GetAllSantas() ([]*SantasDTO, error)
	ChoiceSanta() (*SantasDTO, error)
}

type SantasDTO struct {
	Id       int
	Ime      string
	IsChosen bool
}
