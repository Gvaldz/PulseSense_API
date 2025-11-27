package infrastructure

import (
	"pulse_sense/src/internal/sensores/motion/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type MotionRoutes struct {
	CreateMotionController *controllers.CreateMotionController
	GetByPatientController *controllers.GetByPatientController
}

func NewMotionRoutes(
	createMotionCreateMotionController *controllers.CreateMotionController,
	getBypatientController *controllers.GetByPatientController,
) *MotionRoutes {
	return &MotionRoutes{
		CreateMotionController: createMotionCreateMotionController,
		GetByPatientController: getBypatientController,
	}
}

func (r *MotionRoutes) AttachRoutes(router *gin.Engine) {
	motionsGroup := router.Group("/motions")
	{
		motionsGroup.POST("", r.CreateMotionController.Create)
		motionsGroup.GET("/patient/:IDPatient", r.GetByPatientController.GetByPatient)
	}
}
