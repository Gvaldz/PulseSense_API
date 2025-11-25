package application

import (
	"pulse_sense/src/internal/sensores/patients/domain")

type UpdatePatient struct {
	patientRepo domain.PatientRepository
}

func NewUpdatePatient(patientRepo domain.PatientRepository) *UpdatePatient {
	return &UpdatePatient{patientRepo: patientRepo}
}

func (uc *UpdatePatient) Execute(id string, patient domain.Patient) error {
	return uc.patientRepo.UpdatePatient(id, patient)
}
