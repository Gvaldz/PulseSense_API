package application

import (
	"pulse_sense/src/internal/users/domain")

type GetDoctorsByPatientId struct {
	repo domain.UserRepository
}

func NewGetDoctorsByPatientId(repo domain.UserRepository) *GetDoctorsByPatientId {
	return &GetDoctorsByPatientId{repo: repo}
}

func (cp *GetDoctorsByPatientId) Execute(IDPaciente int32) ([]domain.User, error) {
	return cp.repo.GetDoctorsByPatientId(IDPaciente)
}
