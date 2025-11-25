package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/caregivers/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"

	"github.com/gin-gonic/gin"
)

type CaregiverRoutes struct {
	CreateCaregiversController  	   *controllers.CreateCaregiverController
	CheckCaregiverAssignmentController *controllers.CaregiverAssignmentController
	DeleteCaregiverController          *controllers.DeleteCaregiverController
	TokenService               		   tokenService.TokenService
	AuthRepo                     	   *core.AuthRepository
}

func NewCaregiverRoutes(
	createCaregiversController	 		*controllers.CreateCaregiverController,
	checkCaregiverAssignmentController  *controllers.CaregiverAssignmentController,
	deleteCaregiverController          *controllers.DeleteCaregiverController,
	tokenService 						tokenService.TokenService,
	authRepo							*core.AuthRepository,

) *CaregiverRoutes {
	return &CaregiverRoutes{
		CreateCaregiversController:   		createCaregiversController,
		CheckCaregiverAssignmentController: checkCaregiverAssignmentController,
		DeleteCaregiverController:          deleteCaregiverController,
		TokenService:                 		tokenService,
		AuthRepo:                     		authRepo,
	}
}

func (r *CaregiverRoutes) AttachRoutes(router *gin.Engine) {
    authGroup := router.Group("/")
    {
        authGroup.POST("/caregivers", r.CreateCaregiversController.Create)
		authGroup.GET("/caregivers/assigned/:id", r.CheckCaregiverAssignmentController.CheckAssignment)
		authGroup.DELETE("/caregivers/:idUsuario/:idPaciente", r.DeleteCaregiverController.Delete)
    }
}