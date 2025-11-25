package application

import (
	"pulse_sense/src/internal/caregivers/domain")

type CheckCaregiverAssignment struct {
    repo domain.CaregiverRepository
}

func NewCheckCaregiverAssignment(repo domain.CaregiverRepository) *CheckCaregiverAssignment {
    return &CheckCaregiverAssignment{repo: repo}
}

func (c *CheckCaregiverAssignment) Execute(idUsuario int) (bool, error) {
    return c.repo.IsCaregiverAssigned(idUsuario) 
}