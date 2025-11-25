package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pulse_sense/src/internal/userpatient/application"
)

type GetDoctorByPatientController struct {
	usecase *application.GetDoctorsByPatientId
}

func NewGetDoctorByPatientController(usecase *application.GetDoctorsByPatientId) *GetDoctorByPatientController {
	return &GetDoctorByPatientController{usecase: usecase}
}

func (c *GetDoctorByPatientController) GetByCuidador(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	results, err := c.usecase.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}
