package application

import "pulse_sense/src/internal/userpatient/domain"

type GetDoctorsByPatientId struct {
	repo domain.DoctorPatientRepository
}

func NewGetDoctorsByPatientId(repo domain.DoctorPatientRepository) *GetDoctorsByPatientId {
	return &GetDoctorsByPatientId{repo: repo}
}

func (g *GetDoctorsByPatientId) Execute(idCuidador int32) ([]domain.UserPatientResponse, error) {
	return g.repo.GetDoctorsByPatientId(idCuidador)
}
