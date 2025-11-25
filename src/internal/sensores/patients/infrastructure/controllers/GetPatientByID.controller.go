package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"

	"github.com/gin-gonic/gin"
)

type GetPatientByIDController struct {
	getPatientByID *application.GetPatientByID
}

func NewGetPatientByIDController(getPatientByID *application.GetPatientByID) *GetPatientByIDController {
	return &GetPatientByIDController{getPatientByID: getPatientByID}
}

func (h *GetPatientByIDController) GetPatientByID(c *gin.Context) {
	iduser := c.Param("id")

	user, err := h.getPatientByID.Execute(string(iduser))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
