package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type GetPatientByUser struct {
	repo domain.PatientRepository
}

func NewGetPatientByUser(repo domain.PatientRepository) *GetPatientByUser {
	return &GetPatientByUser{repo: repo}
}

func (cp *GetPatientByUser) Execute(IDPaciente int32) ([]domain.Patient, error) {
	return cp.repo.GetPatientByUser(IDPaciente)
}
