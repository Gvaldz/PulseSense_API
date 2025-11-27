package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/motion/application"
	"pulse_sense/src/internal/sensores/motion/infrastructure/controllers"
	patients "pulse_sense/src/internal/sensores/patients/infrastructure"
	websocket "pulse_sense/src/internal/services/websocket/application"
)

type MotionDependencies struct {
	DB        *sql.DB
	AMQP      *core.AMQPConnection
	WsService *websocket.WebSocketService
	UserRepo  *core.UserRepository
}

func NewMotionDependencies(
	db *sql.DB,
	amqp *core.AMQPConnection,
	wsService *websocket.WebSocketService,
	userRepo *core.UserRepository,
) *MotionDependencies {
	return &MotionDependencies{
		DB:        db,
		AMQP:      amqp,
		WsService: wsService,
		UserRepo:  userRepo,
	}
}

func (d *MotionDependencies) GetRoutes() *MotionRoutes {
	motionRepo := NewMotionRepo(d.DB, nil)
	patientRepo := patients.NewPatientRepo(d.DB)

	createMotionUseCase := application.NewCreateMotion(motionRepo)
	getBypatientUseCase := application.NewGetByPatient(motionRepo)

	createMotionController := controllers.NewCreateMotionController(
		createMotionUseCase,
		d.WsService,
		patientRepo,
		d.UserRepo,
	)
	getBypatientController := controllers.NewGetByPatientController(getBypatientUseCase)

	return NewMotionRoutes(createMotionController, getBypatientController)
}
