package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/motion/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByPatientController struct {
	getByHamster *application.GetByPatient
}

func NewGetByPatientController(getByHamster *application.GetByPatient) *GetByPatientController {
	return &GetByPatientController{getByHamster: getByHamster}
}
func (h *GetByPatientController) GetByPatient(c *gin.Context) {
	idPacienteStr := c.Param("idPaciente")
	idPaciente, err := strconv.Atoi(idPacienteStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid idPaciente"})
		return
	}
	motions, err := h.getByHamster.Execute(idPaciente)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, motions)
}
