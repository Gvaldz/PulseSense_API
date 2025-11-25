package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/signos/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByPatientController struct {
	getByPatient *application.GetByPatient
}

func NewGetByPatientController(getByPatient *application.GetByPatient) *GetByPatientController {
	return &GetByPatientController{getByPatient: getByPatient}
}
func (h *GetByPatientController) GetByPatient(c *gin.Context) {
	IDPacienteStr := c.Param("IDPaciente")
	IDPaciente, err := strconv.Atoi(IDPacienteStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}
	temperatures, err := h.getByPatient.Execute(IDPaciente)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, temperatures)
}
