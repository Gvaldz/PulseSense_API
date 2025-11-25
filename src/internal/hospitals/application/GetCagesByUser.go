package application

import (
	"pulse_sense/src/internal/hospitals/domain")

type GetHospitalByUser struct {
	repo domain.HospitalRepository
}

func NewGetHospitalByUser(repo domain.HospitalRepository) *GetHospitalByUser {
	return &GetHospitalByUser{repo: repo}
}

func (cp *GetHospitalByUser) Execute(idHospital int32) ([]domain.Hospital, error) {
	return cp.repo.GetHospitalByUser(idHospital)
}