package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	patient "pulse_sense/src/internal/sensores/patients/infrastructure"
	"pulse_sense/src/internal/sensores/signos/application"
	"pulse_sense/src/internal/sensores/signos/infrastructure/controllers"
	websocket "pulse_sense/src/internal/services/websocket/application"
)

type SignsDependencies struct {
	DB        *sql.DB
	AMQP      *core.AMQPConnection
	WsService *websocket.WebSocketService
	UserRepo  *core.UserRepository
}

func NewSignsDependencies(
	db *sql.DB,
	amqp *core.AMQPConnection,
	wsService *websocket.WebSocketService,
	userRepo *core.UserRepository,
) *SignsDependencies {
	return &SignsDependencies{
		DB:        db,
		AMQP:      amqp,
		WsService: wsService,
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
	)
	getByPatientController := controllers.NewGetByPatientController(getByPatientUseCase)
	getSignsByTypeController := controllers.NewGetSignsByTypeAndTimeController(getSignsByTypeUseCase)

	return NewSignsRoutes(createSignsController, getByPatientController, getSignsByTypeController,)
}
