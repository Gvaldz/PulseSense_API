package application

import (
	"pulse_sense/src/internal/caregivers/domain")

type CreateCaregiver struct {
	repo domain.CaregiverRepository
}

func NewCreateCaregiver(repo domain.CaregiverRepository) *CreateCaregiver {
	return &CreateCaregiver{repo: repo}
}

func (c *CreateCaregiver) Execute(Caregiver domain.Caregiver) error {
	return c.repo.CreateCaregiver(Caregiver)
}
