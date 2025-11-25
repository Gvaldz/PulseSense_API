package application

import (
	"pulse_sense/src/internal/users/domain"
)

type GetAllUsers struct {
	userRepo domain.UserRepository
}

func NewGetAllUsers(userRepo domain.UserRepository) *GetAllUsers {
	return &GetAllUsers{userRepo: userRepo}
}

func (lp *GetAllUsers) Execute() ([]domain.User, error) {
	return lp.userRepo.GetAllUsers()
}
