package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	patient "pulse_sense/src/internal/sensores/patients/infrastructure"
	"pulse_sense/src/internal/sensores/signos/application"
	"pulse_sense/src/internal/sensores/signos/infrastructure/controllers"
	fcm "pulse_sense/src/internal/services/fcm"
	websocket "pulse_sense/src/internal/services/websocket/application"
)

type SignsDependencies struct {
	DB        *sql.DB
	AMQP      *core.AMQPConnection
	WsService *websocket.WebSocketService
	FCMSender *fcm.FCMSender
	UserRepo  *core.UserRepository
}

func NewSignsDependencies(
	db *sql.DB,
	amqp *core.AMQPConnection,
	wsService *websocket.WebSocketService,
	fcmSender *fcm.FCMSender,
	userRepo *core.UserRepository,
) *SignsDependencies {
	return &SignsDependencies{
		DB:        db,
		AMQP:      amqp,
		WsService: wsService,
		FCMSender: fcmSender,
		UserRepo:  userRepo,
	}
}

func (d *SignsDependencies) GetRoutes() *SignsRoutes {
	signsRepo := NewSignsRepo(d.DB, nil)
	patientRepo := patient.NewPatientRepo(d.DB)

	createSignsUseCase := application.NewCreateSigns(signsRepo)
	getByPatientUseCase := application.NewGetByPatient(signsRepo)
	getSignsByTypeUseCase := application.NewGetSignsByTypeAndTimeRange(signsRepo)

	createSignsController := controllers.NewCreateSignsController(
		createSignsUseCase,
		d.WsService,
		patientRepo,
		d.UserRepo,
		d.FCMSender,
	)
	getByPatientController := controllers.NewGetByPatientController(getByPatientUseCase)
	getSignsByTypeController := controllers.NewGetSignsByTypeAndTimeController(getSignsByTypeUseCase)

	return NewSignsRoutes(createSignsController, getByPatientController, getSignsByTypeController,)
}
