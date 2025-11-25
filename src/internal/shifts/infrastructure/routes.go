package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/shifts/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"
	"github.com/gin-gonic/gin"
)

type ShiftRoutes struct {
	CreateShiftsController      *controllers.CreateShiftController
	TokenService                 tokenService.TokenService
	AuthRepo                     *core.AuthRepository
}

func NewShiftRoutes(
	createShiftsController	 	 *controllers.CreateShiftController,
	tokenService 				 tokenService.TokenService,
	authRepo					 *core.AuthRepository,

) *ShiftRoutes {
	return &ShiftRoutes{
		CreateShiftsController:      createShiftsController,
		TokenService:                 tokenService,
		AuthRepo:                     authRepo,
	}
}

func (r *ShiftRoutes) AttachRoutes(router *gin.Engine) {
    authGroup := router.Group("/")
    {
        authGroup.POST("/shifts", r.CreateShiftsController.Create)
    }
}