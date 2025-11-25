package controllers

import (
	"net/http"
	"pulse_sense/src/internal/caregivers/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCaregiverController struct {
	deleteCaregiver *application.DeleteCaregiver
}

func NewDeleteCaregiverController(delete *application.DeleteCaregiver) *DeleteCaregiverController {
	return &DeleteCaregiverController{
		deleteCaregiver: delete,
	}
}

func (h *DeleteCaregiverController) Delete(c *gin.Context) {
	IdUsuarioStr := c.Param("idUsuario")
	IdPacienteStr := c.Param("idPaciente")

	idUsuario, err := strconv.ParseInt(IdUsuarioStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	IdPaciente, err := strconv.ParseInt(IdPacienteStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de paciente inválido"})
		return
	}

	err = h.deleteCaregiver.Execute(int(idUsuario), int(IdPaciente))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cuidador eliminado correctamente"})
}
