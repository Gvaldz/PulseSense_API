package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"

	"github.com/gin-gonic/gin"
)

type GetAllPatientsController struct {
	getAllPatient *application.GetAllPatients
}

func NewGetAllPatientsController(getAllPatient *application.GetAllPatients) *GetAllPatientsController {
	return &GetAllPatientsController{getAllPatient: getAllPatient}
}

func (h *GetAllPatientsController) GetAllPatient(c *gin.Context) {
	patient, err := h.getAllPatient.Execute()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}
