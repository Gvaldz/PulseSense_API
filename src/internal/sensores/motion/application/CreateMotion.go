package application

import (
	"fmt"
	"pulse_sense/src/internal/sensores/motion/domain"
)

type CreateMotion struct {
	repo domain.MotionRepository
}

func NewCreateMotion(repo domain.MotionRepository) *CreateMotion {
	return &CreateMotion{repo: repo}
}

func (c *CreateMotion) Execute(motion domain.Motion) error {
	fmt.Printf("Guardando movimiento en la base de datos: %+v\n", motion)
	return c.repo.CreateMotion(motion)
}
