package application

import (
	"pulse_sense/src/internal/sensores/signos/domain"
)

type GetByPatient struct {
	repo domain.SignRepository
}

func NewGetByPatient(repo domain.SignRepository) *GetByPatient {
	return &GetByPatient{repo: repo}
}

func (cp *GetByPatient) Execute(IDPaciente int) ([]domain.Sign, error) {
	return cp.repo.GetByPatient(IDPaciente)
}
