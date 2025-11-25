package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type GetHospitalByID struct {
	repo domain.HospitalRepository
}

func NewGetHospitalByID(repo domain.HospitalRepository) *GetHospitalByID {
	return &GetHospitalByID{repo: repo}
}

func (cp *GetHospitalByID) Execute(IDRaza string) (domain.Hospital, error) {
	return cp.repo.GetHospitalByID(IDRaza)
}
