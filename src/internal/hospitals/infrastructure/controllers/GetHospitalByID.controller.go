package controllers

import (
	"net/http"
	"pulse_sense/src/internal/hospitals/application"

	"github.com/gin-gonic/gin"
)

type GetHospitalByIDController struct {
	getHospitalByID *application.GetHospitalByID
}

func NewGetHospitalByIDController(getHospitalByID *application.GetHospitalByID) *GetHospitalByIDController {
	return &GetHospitalByIDController{getHospitalByID: getHospitalByID}
}

func (h *GetHospitalByIDController) GetHospitalByID(c *gin.Context) {
	iduser := c.Param("id")

	user, err := h.getHospitalByID.Execute(string(iduser))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
