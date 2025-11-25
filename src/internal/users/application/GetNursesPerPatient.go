package application

import (
	"pulse_sense/src/internal/users/domain"
)

type GetNursePerPatient struct {
	userRepo domain.UserRepository
}

func NewGetNursePerPatient(userRepo domain.UserRepository) *GetNursePerPatient {
	return &GetNursePerPatient{userRepo: userRepo}
}

func (lp *GetNursePerPatient) Execute(idpatient int32) ([]domain.User, error) {
	return lp.userRepo.GetNursePerPatient(idpatient)
}
