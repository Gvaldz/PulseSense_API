package application

import (
	"fmt"
	"pulse_sense/src/internal/sensores/signos/domain"
)

type CreateSigns struct {
	repo domain.SignRepository
}

func NewCreateSigns(repo domain.SignRepository) *CreateSigns {
	return &CreateSigns{repo: repo}
}

func (c *CreateSigns) Execute(sign domain.Sign) error {
	fmt.Printf("Guardando signo en la base de datos: %+v\n", sign)
	return c.repo.CreateSigns(sign)
}
