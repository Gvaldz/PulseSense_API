package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type GetAllPatients struct {
	repo domain.PatientRepository
}

func NewGetAllPatients(repo domain.PatientRepository) *GetAllPatients {
	return &GetAllPatients{repo: repo}
}

func (cp *GetAllPatients) Execute() ([]domain.Patient, error) {
	return cp.repo.GetAllPatient()
}
