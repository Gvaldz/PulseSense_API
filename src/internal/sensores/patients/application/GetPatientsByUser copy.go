package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type GetPatientByNurse struct {
	repo domain.PatientRepository
}

func NewGetPatientByNurse(repo domain.PatientRepository) *GetPatientByNurse {
	return &GetPatientByNurse{repo: repo}
}

func (cp *GetPatientByNurse) Execute(IDPaciente int32) ([]domain.Patient, error) {
	return cp.repo.GetPatientByNurse(IDPaciente)
}
