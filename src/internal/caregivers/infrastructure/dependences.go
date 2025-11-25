package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/caregivers/application"
	"pulse_sense/src/internal/caregivers/infrastructure/controllers"
)

type CaregiverDependencies struct {
	DB *sql.DB
}

func NewCaregiverDependencies(db *sql.DB) *CaregiverDependencies {
	return &CaregiverDependencies{DB: db}
}

func (d *CaregiverDependencies) GetRoutes() *CaregiverRoutes {
	CaregiverRepo := NewCaregiverRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	createCaregiverUseCase := application.NewCreateCaregiver(CaregiverRepo)
	createCaregiverController := controllers.NewCreateCaregiverController(createCaregiverUseCase)
	deleteCaregiverUseCase := application.NewDeleteCaregiver(CaregiverRepo)
	deleteCaregiverController := controllers.NewDeleteCaregiverController(deleteCaregiverUseCase)
	checkCaregiverAssignmentUseCase := application.NewCheckCaregiverAssignment(CaregiverRepo)
	checkCaregiverAssignmentController := controllers.NewCaregiverAssignmentController(checkCaregiverAssignmentUseCase)

	return NewCaregiverRoutes(createCaregiverController, checkCaregiverAssignmentController, deleteCaregiverController, tokenService, authRepo)
}
