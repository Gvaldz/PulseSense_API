package application

import (
	"pulse_sense/src/internal/caregivers/domain"
)

type DeleteCaregiver struct {
	repo domain.CaregiverRepository
}

func NewDeleteCaregiver(repo domain.CaregiverRepository) *DeleteCaregiver {
	return &DeleteCaregiver{repo: repo}
}

func (lp *DeleteCaregiver) Execute(IdUsuario int, IdPaciente int) error {
	return lp.repo.DeleteCaregiver(IdUsuario, IdPaciente)
}
