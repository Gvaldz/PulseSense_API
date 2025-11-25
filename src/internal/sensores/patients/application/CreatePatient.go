package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type CreatePatient struct {
    repo domain.PatientRepository
}

func NewCreatePatient(repo domain.PatientRepository) *CreatePatient {
    return &CreatePatient{repo: repo}
}

func (c *CreatePatient) Execute(patient domain.Patient) (int64, error) {
    return c.repo.CreatePatient(patient)
}