package controllers

import (
	"net/http"
	"pulse_sense/src/internal/hospitals/application"
	"pulse_sense/src/internal/hospitals/domain"
	"github.com/gin-gonic/gin"
)

type UpdateHospitalController struct {
	updateHospitalController *application.UpdateHospital
}

func NewUpdateHospitalController(updateHospitalController *application.UpdateHospital) *UpdateHospitalController {
	return &UpdateHospitalController{
		updateHospitalController: updateHospitalController,
	}
}

func (c *UpdateHospitalController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user domain.Hospital
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.updateHospitalController.Execute(id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Hospital actualizado correctamente"})
}
