package application

import (
	"pulse_sense/src/internal/users/domain"
)

type GetNursePerHospital struct {
	userRepo domain.UserRepository
}

func NewGetNursePerHospital(userRepo domain.UserRepository) *GetNursePerHospital {
	return &GetNursePerHospital{userRepo: userRepo}
}

func (lp *GetNursePerHospital) Execute(idHospital int32) ([]domain.User, error) {
	return lp.userRepo.GetNursePerHospital(idHospital)
}
