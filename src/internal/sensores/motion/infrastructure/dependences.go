package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/motion/application"
	"pulse_sense/src/internal/sensores/motion/infrastructure/controllers"
	patients "pulse_sense/src/internal/sensores/patients/infrastructure"
	fcm "pulse_sense/src/internal/services/fcm"
	websocket "pulse_sense/src/internal/services/websocket/application"
)

type MotionDependencies struct {
	DB        *sql.DB
	AMQP      *core.AMQPConnection
	WsService *websocket.WebSocketService
	FCMSender *fcm.FCMSender
	UserRepo  *core.UserRepository
}

func NewMotionDependencies(
	db *sql.DB,
	amqp *core.AMQPConnection,
	wsService *websocket.WebSocketService,
	fcmSender *fcm.FCMSender,
	userRepo *core.UserRepository,
) *MotionDependencies {
	return &MotionDependencies{
		DB:        db,
		AMQP:      amqp,
		WsService: wsService,
		FCMSender: fcmSender,
		UserRepo:  userRepo,
	}
}

func (d *MotionDependencies) GetRoutes() *MotionRoutes {
	motionRepo := NewMotionRepo(d.DB, nil)
	cageRepo := patients.NewPatientRepo(d.DB)

	createMotionUseCase := application.NewCreateMotion(motionRepo)
	getByHamsterUseCase := application.NewGetByPatient(motionRepo)

	createMotionController := controllers.NewCreateMotionController(
		createMotionUseCase,
		d.WsService,
		cageRepo,
		d.UserRepo,
		d.FCMSender,
	)
	getByHamsterController := controllers.NewGetByPatientController(getByHamsterUseCase)

	return NewMotionRoutes(createMotionController, getByHamsterController)
}
