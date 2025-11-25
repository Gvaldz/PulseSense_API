package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type UpdateHospital struct {
	HospitalRepo domain.HospitalRepository
}

func NewUpdateHospital(HospitalRepo domain.HospitalRepository) *UpdateHospital {
	return &UpdateHospital{HospitalRepo: HospitalRepo}
}

func (uc *UpdateHospital) Execute(id string, Hospital domain.Hospital) error {
	return uc.HospitalRepo.UpdateHospital(id, Hospital)
}
