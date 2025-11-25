package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/patients/application"
	"pulse_sense/src/internal/sensores/patients/infrastructure/controllers"
)

type PatientDependencies struct {
	DB *sql.DB
}

func NewPatientDependencies(db *sql.DB) *PatientDependencies {
	return &PatientDependencies{DB: db}
}

func (d *PatientDependencies) GetRoutes() *PatientRoutes {
	PatientRepo := NewPatientRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	createPatientUseCase := application.NewCreatePatient(PatientRepo)
	getPatientUseCase := application.NewGetPatientByID(PatientRepo)
	getPatientByUserUseCase := application.NewGetPatientByUser(PatientRepo)
	getPatientByNurseUseCase := application.NewGetPatientByNurse(PatientRepo)
	updatePatientUseCase := application.NewUpdatePatient(PatientRepo)

	createPatientController := controllers.NewCreatePatientController(createPatientUseCase)
	getPatientByIdController := controllers.NewGetPatientByIDController(getPatientUseCase)
	getPatientByUserController := controllers.NewGetPatientByUserController(getPatientByUserUseCase)
	getPatientByNurseController := controllers.NewGetPatientByNurseController(getPatientByNurseUseCase)
	updatePatientController := controllers.NewUpdatePatientController(updatePatientUseCase)

	return NewPatientRoutes(createPatientController, getPatientByIdController, getPatientByUserController, getPatientByNurseController, updatePatientController, tokenService, authRepo)
}
