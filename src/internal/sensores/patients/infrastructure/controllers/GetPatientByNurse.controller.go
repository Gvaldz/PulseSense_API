package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPatientByNurseController struct {
	getByNurse *application.GetPatientByNurse
}

func NewGetPatientByNurseController(getByNurse *application.GetPatientByNurse) *GetPatientByNurseController {
	return &GetPatientByNurseController{getByNurse: getByNurse}
}

func (h *GetPatientByNurseController) GetByNurse(c *gin.Context) {
	idNurse := c.Param("id")
	idInt, err := strconv.Atoi(idNurse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	Nurse, err := h.getByNurse.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Nurse,
	})
}
