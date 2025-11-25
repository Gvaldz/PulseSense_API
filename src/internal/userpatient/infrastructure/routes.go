package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/userpatient/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"
	"github.com/gin-gonic/gin"
)

type UserPatientRoutes struct {
	GetDoctorsByPatientId       *controllers.GetDoctorByPatientController
	GetNurseByPatientId         *controllers.GetNurseByPatientController
	TokenService                 tokenService.TokenService
	AuthRepo                     *core.AuthRepository
}

func NewUserPatientRoutes(
	getDoctorsByPatientId	 	 *controllers.GetDoctorByPatientController,
	getNurseByPatientId 		 *controllers.GetNurseByPatientController,
	tokenService 				 tokenService.TokenService,
	authRepo					 *core.AuthRepository,

) *UserPatientRoutes {
	return &UserPatientRoutes{
		GetDoctorsByPatientId:        getDoctorsByPatientId,
		GetNurseByPatientId:          getNurseByPatientId,
		TokenService:                 tokenService,
		AuthRepo:                     authRepo,
	}
}

func (r *UserPatientRoutes) AttachRoutes(router *gin.Engine) {
    authGroup := router.Group("/")
    {
        authGroup.GET("/doctorpatient/:id", r.GetDoctorsByPatientId.GetByCuidador)
		authGroup.GET("/nursepatient/:id", r.GetNurseByPatientId.GetByCuidador)
    }
}