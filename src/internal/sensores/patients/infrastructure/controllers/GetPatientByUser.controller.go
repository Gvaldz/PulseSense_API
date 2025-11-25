package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetPatientByUserController struct {
	getByUser *application.GetPatientByUser
}

func NewGetPatientByUserController(getByUser *application.GetPatientByUser) *GetPatientByUserController {
	return &GetPatientByUserController{getByUser: getByUser}
}

func (h *GetPatientByUserController) GetByUser(c *gin.Context) {
	iduser := c.Param("id")
	idInt, err := strconv.Atoi(iduser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	user, err := h.getByUser.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
