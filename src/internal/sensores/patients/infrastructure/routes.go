package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/patients/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"
	"pulse_sense/src/server/middleware"

	"github.com/gin-gonic/gin"
)

type PatientRoutes struct {
	CreatePatientController    	*controllers.CreatePatientController
	GetPatientController       	*controllers.GetPatientByIDController
	GetPatientByUserController 	*controllers.GetPatientByUserController
	GetPatientByNurseController	*controllers.GetPatientByNurseController
	UpdatePatientController     *controllers.UpdatePatientController
	TokenService                tokenService.TokenService
	AuthRepo                    *core.AuthRepository
}

func NewPatientRoutes(
	createPatientController 	*controllers.CreatePatientController,
	getPatientController 		*controllers.GetPatientByIDController,
	getPatientByUserController 	*controllers.GetPatientByUserController,
	getPatientByNurseController	*controllers.GetPatientByNurseController,
	updatePatientController 	*controllers.UpdatePatientController,
	tokenService 				tokenService.TokenService,
	authRepo 					*core.AuthRepository,

) *PatientRoutes {
	return &PatientRoutes{
		CreatePatientController:    createPatientController,
		GetPatientController:       getPatientController,
		GetPatientByUserController: getPatientByUserController,
		GetPatientByNurseController: getPatientByNurseController,
		UpdatePatientController:    updatePatientController,
		TokenService:               tokenService,
		AuthRepo:                   authRepo,
	}
}

func (r *PatientRoutes) AttachRoutes(router *gin.Engine) {
	doctorAuth := middleware.AuthMiddleware(r.TokenService, r.AuthRepo, "1")
	nurseAuth := middleware.AuthMiddleware(r.TokenService, r.AuthRepo, "2" )

	nurseGroup := router.Group("/nurse/patient")
	nurseGroup.Use(nurseAuth)
	{
		nurseGroup.GET("/:id", r.GetPatientController.GetPatientByID)
		nurseGroup.GET("/user/:id", r.GetPatientByNurseController.GetByNurse)
	}

	doctorGroup := router.Group("/doctor/patient")
	doctorGroup.Use(doctorAuth)
	{
		doctorGroup.POST("", r.CreatePatientController.Create)
		doctorGroup.PUT("/:id", r.UpdatePatientController.UpdateUser)
		doctorGroup.GET("/:id", r.GetPatientController.GetPatientByID)
		doctorGroup.GET("/user/:id", r.GetPatientByUserController.GetByUser)
	}
}
