package application

import (
	"pulse_sense/src/internal/sensores/motion/domain"
)

type GetByPatient struct {
	repo domain.MotionRepository
}

func NewGetByPatient(repo domain.MotionRepository) *GetByPatient {
	return &GetByPatient{repo: repo}
}

func (cp *GetByPatient) Execute(IDPatient int) ([]domain.Motion, error) {
	return cp.repo.GetByPatient(IDPatient)
}
