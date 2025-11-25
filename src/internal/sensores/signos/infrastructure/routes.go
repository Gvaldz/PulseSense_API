package infrastructure

import (
	"pulse_sense/src/internal/sensores/signos/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type SignsRoutes struct {
	CreateSignsController 		*controllers.CreateSignsController
	GetByPatientController      *controllers.GetByPatientController
	GetSignsByTypeController    *controllers.GetSignsByTypeAndTimeController
}

func NewSignsRoutes(
	createSignsController		*controllers.CreateSignsController,
	getByPatientController 		*controllers.GetByPatientController,
	getSignsByTypeController 	*controllers.GetSignsByTypeAndTimeController,
) *SignsRoutes {
	return &SignsRoutes{
		CreateSignsController: createSignsController,
		GetByPatientController:      getByPatientController,
		GetSignsByTypeController:    getSignsByTypeController,
	}
}

func (r *SignsRoutes) AttachRoutes(router *gin.Engine) {
    SignsGroup := router.Group("/signs")
    {
        SignsGroup.POST("", r.CreateSignsController.Create)
        SignsGroup.GET("/patient/:IDPaciente", r.GetByPatientController.GetByPatient)
        SignsGroup.GET("/patient/:IDPaciente/:IDTipo/:fecha/:turno", r.GetSignsByTypeController.GetByTypeAndTime)
    }
}