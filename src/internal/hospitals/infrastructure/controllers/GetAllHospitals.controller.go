package controllers

import (
	"net/http"
	"pulse_sense/src/internal/hospitals/application"

	"github.com/gin-gonic/gin"
)

type GetAllHospitalsController struct {
	getAllHospital *application.GetAllHospitals
}

func NewGetAllHospitalsController(getAllHospital *application.GetAllHospitals) *GetAllHospitalsController {
	return &GetAllHospitalsController{getAllHospital: getAllHospital}
}

func (h *GetAllHospitalsController) GetAllHospital(c *gin.Context) {
	Hospital, err := h.getAllHospital.Execute()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Hospital)
}
