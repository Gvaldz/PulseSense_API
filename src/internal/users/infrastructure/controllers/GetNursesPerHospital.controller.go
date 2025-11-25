package controllers

import (
	"net/http"
	"pulse_sense/src/internal/users/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetNursePerHospitalController struct {
	NursePerHospitalService *application.GetNursePerHospital
}

func NewGetNursePerHospitalController(getNursePerHospital *application.GetNursePerHospital) *GetNursePerHospitalController {
	return &GetNursePerHospitalController{NursePerHospitalService: getNursePerHospital}
}

func (h *GetNursePerHospitalController) GetNursePerHospital(c *gin.Context) {
	idhospital := c.Param("id")
	idInt, err := strconv.Atoi(idhospital)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de hospital inválido"})
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
