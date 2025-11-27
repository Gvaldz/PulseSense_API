package infrastructure

import (
	"database/sql"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/users/application"
	"pulse_sense/src/internal/users/infrastructure/controllers"
	middleware "pulse_sense/src/server/middleware"
)

type UserDependencies struct {
	DB           *sql.DB
	AMQP         *core.AMQPConnection
	Hasher       *core.BcryptHasher
	UserRepo     *core.UserRepository
	AuthRepo     *core.AuthRepository
	TokenService *core.JWTService
}

func NewUserDependencies(db *sql.DB, amqp *core.AMQPConnection, hasher *core.BcryptHasher, tokenService *core.JWTService, authRepo *core.AuthRepository, userRepo *core.UserRepository) *UserDependencies {

	return &UserDependencies{
		DB:           db,
		AMQP:         amqp,
		Hasher:       hasher,
		TokenService: tokenService,
		AuthRepo:     authRepo,
		UserRepo:     userRepo,
	}
}

func (d *UserDependencies) GetRoutes() *UserRoutes {
	createUserUseCase := application.NewCreateUser(d.UserRepo, d.Hasher)
	getAllUserUseCase := application.NewGetAllUsers(d.UserRepo)
	getUserUseCase := application.NewGetUserByID(d.UserRepo)
	updateUserUseCase := application.NewUpdateUser(d.UserRepo)
	updatePasswordUseCase := application.NewUpdatePassword(d.UserRepo, d.Hasher)
	deleteUserUseCase := application.NewDeleteUser(d.UserRepo)
	getDoctorsByPatient := application.NewGetDoctorsByPatientId(d.UserRepo)
	getNursePerHospitalUseCase := application.NewGetNursePerHospital(d.UserRepo)
	getNursePerPatientUseCase := application.NewGetNursePerPatient(d.UserRepo)
	fcmController := controllers.NewFCMController(d.UserRepo)

	createUserController := controllers.NewCreateUserController(createUserUseCase)
	getUsersController := controllers.NewGetAllUsersController(getAllUserUseCase)
	getUserController := controllers.NewGetByUserIDController(getUserUseCase)
	updateUserController := controllers.NewUpdateUserController(updateUserUseCase)
	updatePasswordController := controllers.NewUpdatePasswordController(updatePasswordUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)
	getDoctorsByPatientController := controllers.NewGetDoctorsByPatientIdController(getDoctorsByPatient)
	getNursePerHospitalController := controllers.NewGetNursePerHospitalController(getNursePerHospitalUseCase)
	getNursePerPatientController := controllers.NewGetNursePerPatientController(getNursePerPatientUseCase)
	authMiddleware := middleware.AuthMiddleware(d.TokenService, d.AuthRepo, "usuario")

	return NewUserRoutes(
		createUserController,
		getUsersController,
		getUserController,
		updateUserController,
		updatePasswordController,
		deleteUserController,
		getDoctorsByPatientController,
		getNursePerHospitalController,
		getNursePerPatientController,
		fcmController,
		authMiddleware,
	)
}
