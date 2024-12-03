package service

import "blazperic/santa/port"

type SantaService struct {
	santaService port.Database
}

func NewSantaService(santaService port.Database) *SantaService {
	return &SantaService{
		santaService: santaService,
	}
}

func (s *SantaService) GetAllSantas() ([]*port.SantasDTO, error) {
	return s.santaService.GetAllSantas()
}

func (s *SantaService) ChoiceSanta() (*port.SantasDTO, error) {
	return s.santaService.ChoiceSanta()
}
