package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/userpatient/application"
	"pulse_sense/src/internal/userpatient/infrastructure/controllers"
)

type UserPatientDependencies struct {
	DB *sql.DB
}

func NewUserPatientDependencies(db *sql.DB) *UserPatientDependencies {
	return &UserPatientDependencies{DB: db}
}

func (d *UserPatientDependencies) GetRoutes() *UserPatientRoutes {
	UserPatientRepo := NewUserPatientRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	getDoctorByPatientUseCase := application.NewGetDoctorsByPatientId(UserPatientRepo)
	getDosctorByPatientController := controllers.NewGetDoctorByPatientController(getDoctorByPatientUseCase)

	getNurseByPatientUseCase := application.NewGetNursesByPatientId(UserPatientRepo)
	getNurseByPatientController := controllers.NewGetNurseByPatientController(getNurseByPatientUseCase)

	return NewUserPatientRoutes(getDosctorByPatientController, getNurseByPatientController ,tokenService, authRepo)
}
