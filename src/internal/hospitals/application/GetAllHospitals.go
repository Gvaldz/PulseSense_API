package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type GetAllHospitals struct {
	repo domain.HospitalRepository
}

func NewGetAllHospitals(repo domain.HospitalRepository) *GetAllHospitals {
	return &GetAllHospitals{repo: repo}
}

func (cp *GetAllHospitals) Execute() ([]domain.Hospital, error) {
	return cp.repo.GetAllHospital()
}
