package app

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/server"
	consumer_amqp 	"pulse_sense/src/consumer_amqp"
	patient 		"pulse_sense/src/internal/sensores/patients/infrastructure"
	sign 			"pulse_sense/src/internal/sensores/signos/infrastructure"
	login 			"pulse_sense/src/internal/services/auth/infrastructure"
	websocketapp 	"pulse_sense/src/internal/services/websocket/application"
	websocketinfra  "pulse_sense/src/internal/services/websocket/infrastructure"
	websocketc 		"pulse_sense/src/internal/services/websocket/infrastructure/controllers"
	users 			"pulse_sense/src/internal/users/infrastructure"
	hospitals 		"pulse_sense/src/internal/hospitals/infrastructure"
	workers 		"pulse_sense/src/internal/workers/infrastructure"
	caregivers 		"pulse_sense/src/internal/caregivers/infrastructure"
	shifts 			"pulse_sense/src/internal/shifts/infrastructure"
	motion 			"pulse_sense/src/internal/sensores/motion/infrastructure"
	userpatient 	"pulse_sense/src/internal/userpatient/infrastructure"
	middleware 		"pulse_sense/src/server/middleware"
)

type Application struct {
	DB           *sql.DB
	AMQPConn     *core.AMQPConnection
	Server       *server.Server
	AMQPConsumer *consumer_amqp.RabbitMQConsumer
	Hasher       *core.BcryptHasher
	tokenService *core.JWTService
}

func NewApplication() (*Application, error) {
	db, err := core.ConnectDB()
	if err != nil {
		return nil, err
	}

	amqpConn, err := core.NewAMQPConnection()


	hasher := core.NewBcryptHasher(12)
	wsService := websocketapp.NewWebSocketService()
	tokenService := core.NewJWTService()
	userRepo := core.NewUserRepository(db).(*core.UserRepository)
	authRepo := &core.AuthRepository{DB: db}

	usersDeps := users.NewUserDependencies(
		db,
		amqpConn,
		hasher,
		tokenService,
		authRepo,
		userRepo,
	)

	loginDeps := login.NewAuthDependencies(
		db,
		hasher,
		userRepo,
	)

	authMiddleware := middleware.AuthMiddleware(
		tokenService,
		authRepo,
		"usuario",
	)

	wsHandler := websocketc.NewWebSocketController(wsService, *tokenService)
	wsRoutes := websocketinfra.NewWebSocketRoutes(wsHandler)
	patientDeps := patient.NewPatientDependencies(db)
	signDeps := sign.NewSignsDependencies(db, amqpConn, wsService, userRepo)
	hospitalDeps := hospitals.NewHospitalDependencies(db)
	workerDeps := workers.NewWorkerDependencies(db)
	caregiversDeps := caregivers.NewCaregiverDependencies(db)
	shiftsDeps := shifts.NewShiftDependencies(db)
	userpatientDeps := userpatient.NewUserPatientDependencies(db)
	motionDeps := motion.NewMotionDependencies(db, amqpConn, wsService, userRepo)

	server := server.NewServer(
		signDeps.GetRoutes(),
		usersDeps.GetRoutes(),
		patientDeps.GetRoutes(),
		loginDeps.GetRoutes(),
		hospitalDeps.GetRoutes(),
		workerDeps.GetRoutes(),
		caregiversDeps.GetRoutes(),
		userpatientDeps.GetRoutes(),
		shiftsDeps.GetRoutes(),
		wsRoutes,
		authMiddleware,
	)

	consumer := consumer_amqp.NewRabbitMQConsumer(
		amqpConn,
		signDeps.GetRoutes().CreateSignsController,
		patientDeps.GetRoutes().CreatePatientController,
		motionDeps.GetRoutes().CreateMotionController,
	)

	return &Application{
		DB:           db,
		AMQPConn:     amqpConn,
		Server:       server,
		AMQPConsumer: consumer,
		Hasher:       hasher,
		tokenService: tokenService,
	}, nil
}

func (a *Application) Start() error {
	go a.AMQPConsumer.Start()
	return a.Server.Run()
}

func (a *Application) Close() {
	if a.AMQPConn != nil {
		a.AMQPConn.Close()
	}
}
