package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/hospitals/application"
	"pulse_sense/src/internal/hospitals/infrastructure/controllers"
)

type HospitalDependencies struct {
	DB *sql.DB
}

func NewHospitalDependencies(db *sql.DB) *HospitalDependencies {
	return &HospitalDependencies{DB: db}
}

func (d *HospitalDependencies) GetRoutes() *HospitalRoutes {
	HospitalRepo := NewHospitalRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	createHospitalUseCase := application.NewCreateHospital(HospitalRepo)
	getAllHospitalUseCase := application.NewGetAllHospitals(HospitalRepo)
	getHospitalUseCase := application.NewGetHospitalByID(HospitalRepo)
	getHospitalByUserUseCase := application.NewGetHospitalByUser(HospitalRepo)
	updateHospitalUseCase := application.NewUpdateHospital(HospitalRepo)
	SearchHospitalUseCase := application.NewSearchHospital(HospitalRepo)

	createHospitalController := controllers.NewCreateHospitalController(createHospitalUseCase)
	getAllHospitalController := controllers.NewGetAllHospitalsController(getAllHospitalUseCase)
	getHospitalByIdController := controllers.NewGetHospitalByIDController(getHospitalUseCase)
	getHospitalByUserController := controllers.NewGetHospitalByUserController(getHospitalByUserUseCase)
	updateHospitalController := controllers.NewUpdateHospitalController(updateHospitalUseCase)
	searchHospitalController := controllers.NewSearchHospitalController(SearchHospitalUseCase)

	return NewHospitalRoutes(createHospitalController, getAllHospitalController, getHospitalByIdController, getHospitalByUserController, updateHospitalController, searchHospitalController, tokenService, authRepo)
}
