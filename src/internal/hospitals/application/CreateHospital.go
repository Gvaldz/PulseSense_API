package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type CreateHospital struct {
	repo domain.HospitalRepository
}

func NewCreateHospital(repo domain.HospitalRepository) *CreateHospital {
	return &CreateHospital{repo: repo}
}

func (c *CreateHospital) Execute(Hospital domain.Hospital) error {
	return c.repo.CreateHospital(Hospital)
}
