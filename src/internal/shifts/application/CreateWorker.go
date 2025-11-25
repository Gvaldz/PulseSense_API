package application

import (
	"pulse_sense/src/internal/shifts/domain")

type CreateShift struct {
	repo domain.ShiftRepository
}

func NewCreateShift(repo domain.ShiftRepository) *CreateShift {
	return &CreateShift{repo: repo}
}

func (c *CreateShift) Execute(Shift domain.Shift) error {
	return c.repo.CreateShift(Shift)
}
