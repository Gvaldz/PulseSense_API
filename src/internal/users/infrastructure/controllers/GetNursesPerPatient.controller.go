package controllers

import (
	"net/http"
	"pulse_sense/src/internal/users/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetNursePerPatientController struct {
	NursePerHospitalService *application.GetNursePerPatient
}

func NewGetNursePerPatientController(getNursePerPatient *application.GetNursePerPatient) *GetNursePerPatientController {
	return &GetNursePerPatientController{NursePerHospitalService: getNursePerPatient}
}

func (h *GetNursePerPatientController) GetNursePerPatient(c *gin.Context) {
	idpatient := c.Param("id")
	idInt, err := strconv.Atoi(idpatient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de paciente inválido"})
		return
	}

	user, err := h.NursePerHospitalService.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
