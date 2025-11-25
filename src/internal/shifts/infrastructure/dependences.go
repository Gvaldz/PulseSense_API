package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/shifts/application"
	"pulse_sense/src/internal/shifts/infrastructure/controllers"
)

type ShiftDependencies struct {
	DB *sql.DB
}

func NewShiftDependencies(db *sql.DB) *ShiftDependencies {
	return &ShiftDependencies{DB: db}
}

func (d *ShiftDependencies) GetRoutes() *ShiftRoutes {
	ShiftRepo := NewShiftRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	createShiftUseCase := application.NewCreateShift(ShiftRepo)
	createShiftController := controllers.NewCreateShiftController(createShiftUseCase)

	return NewShiftRoutes(createShiftController, tokenService, authRepo)
}
