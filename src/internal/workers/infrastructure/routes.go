package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/workers/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"

	"github.com/gin-gonic/gin"
)

type WorkerRoutes struct {
	CreateWorkersController      *controllers.CreateWorkerController
	CheckWorkerAssignmentController *controllers.WorkerAssignmentController
	TokenService                 tokenService.TokenService
	AuthRepo                     *core.AuthRepository
}

func NewWorkerRoutes(
	createWorkersController	 	 *controllers.CreateWorkerController,
	checkWorkerAssignmentController *controllers.WorkerAssignmentController,
	tokenService 				 tokenService.TokenService,
	authRepo					 *core.AuthRepository,

) *WorkerRoutes {
	return &WorkerRoutes{
		CreateWorkersController:      createWorkersController,
		CheckWorkerAssignmentController: checkWorkerAssignmentController,
		TokenService:                 tokenService,
		AuthRepo:                     authRepo,
	}
}

func (r *WorkerRoutes) AttachRoutes(router *gin.Engine) {
    authGroup := router.Group("/")
    {
        authGroup.POST("/workers", r.CreateWorkersController.Create)
		authGroup.GET("/workers/verify/:id", r.CheckWorkerAssignmentController.CheckAssignment)
    }
}