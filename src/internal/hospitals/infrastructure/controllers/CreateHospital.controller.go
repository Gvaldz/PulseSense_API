package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/hospitals/application"
	"pulse_sense/src/internal/hospitals/domain"

	"github.com/gin-gonic/gin"
)

type CreateHospitalController struct {
	CreateHospital *application.CreateHospital
}

func NewCreateHospitalController(CreateHospital *application.CreateHospital) *CreateHospitalController {
	return &CreateHospitalController{
		CreateHospital: CreateHospital,
	}
}

func (h *CreateHospitalController) Create(c *gin.Context) {
	var HospitalRequest domain.Hospital
	if err := c.ShouldBindJSON(&HospitalRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CreateHospital.Execute(HospitalRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Hospital creado correctamente", "Hospital": HospitalRequest})
}

func (h *CreateHospitalController) ProcessHospital(Hospital domain.Hospital) error {
	fmt.Printf("Procesando creación de hospital: %+v\n", Hospital)
	return h.CreateHospital.Execute(Hospital)
}
