package application

import (
	"pulse_sense/src/internal/sensores/signos/domain"
)

type GetSignsByTypeAndTimeRange struct {
	repo domain.SignRepository
}

func NewGetSignsByTypeAndTimeRange(repo domain.SignRepository) *GetSignsByTypeAndTimeRange {
	return &GetSignsByTypeAndTimeRange{repo: repo}
}

func (cp *GetSignsByTypeAndTimeRange) Execute(IDPaciente int, IDTipo int, fecha string, turno string) ([]domain.Sign, error) {
	return cp.repo.GetSignsByTypeAndTimeRange(IDPaciente, IDTipo, fecha, turno)
}
