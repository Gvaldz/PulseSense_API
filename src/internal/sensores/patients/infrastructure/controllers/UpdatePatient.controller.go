package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"
	"pulse_sense/src/internal/sensores/patients/domain"
	"github.com/gin-gonic/gin"
)

type UpdatePatientController struct {
	updatePatientController *application.UpdatePatient
}

func NewUpdatePatientController(updatePatientController *application.UpdatePatient) *UpdatePatientController {
	return &UpdatePatientController{
		updatePatientController: updatePatientController,
	}
}

func (c *UpdatePatientController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user domain.Patient
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.updatePatientController.Execute(id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Paciente actualizado correctamente"})
}
