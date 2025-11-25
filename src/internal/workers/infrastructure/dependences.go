package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/workers/application"
	"pulse_sense/src/internal/workers/infrastructure/controllers"
)

type WorkerDependencies struct {
	DB *sql.DB
}

func NewWorkerDependencies(db *sql.DB) *WorkerDependencies {
	return &WorkerDependencies{DB: db}
}

func (d *WorkerDependencies) GetRoutes() *WorkerRoutes {
	WorkerRepo := NewWorkerRepo(d.DB)
	tokenService := core.NewJWTService()
	authRepo := core.NewAuthRepository(d.DB).(*core.AuthRepository)

	createWorkerUseCase := application.NewCreateWorker(WorkerRepo)
	checkWorkerAssignmentUseCase := application.NewCheckWorkerAssignment(WorkerRepo)

	createWorkerController := controllers.NewCreateWorkerController(createWorkerUseCase)
	checkWorkerAssignmentController := controllers.NewWorkerAssignmentController(checkWorkerAssignmentUseCase)

	return NewWorkerRoutes(createWorkerController, checkWorkerAssignmentController, tokenService, authRepo)
}
