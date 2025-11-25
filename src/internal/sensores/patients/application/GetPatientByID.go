package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type GetPatientByID struct {
	repo domain.PatientRepository
}

func NewGetPatientByID(repo domain.PatientRepository) *GetPatientByID {
	return &GetPatientByID{repo: repo}
}

func (cp *GetPatientByID) Execute(IDRaza string) (domain.Patient, error) {
	return cp.repo.GetPatientByID(IDRaza)
}
