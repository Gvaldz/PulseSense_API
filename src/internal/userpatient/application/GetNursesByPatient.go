package application

import "pulse_sense/src/internal/userpatient/domain"

type GetNursesByPatientId struct {
	repo domain.DoctorPatientRepository
}

func NewGetNursesByPatientId(repo domain.DoctorPatientRepository) *GetNursesByPatientId {
	return &GetNursesByPatientId{repo: repo}
}

func (g *GetNursesByPatientId) Execute(idCuidador int32) ([]domain.UserPatientResponse, error) {
	return g.repo.GetNursesByPatientId(idCuidador)
}
