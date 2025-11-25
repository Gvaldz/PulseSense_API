package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/caregivers/application"
	"pulse_sense/src/internal/caregivers/domain"

	"github.com/gin-gonic/gin"
)

type CreateCaregiverController struct {
	CreateCaregiver *application.CreateCaregiver
}

func NewCreateCaregiverController(CreateCaregiver *application.CreateCaregiver) *CreateCaregiverController {
	return &CreateCaregiverController{
		CreateCaregiver: CreateCaregiver,
	}
}

func (h *CreateCaregiverController) Create(c *gin.Context) {
	var CaregiverRequest domain.Caregiver
	if err := c.ShouldBindJSON(&CaregiverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CreateCaregiver.Execute(CaregiverRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "cuidador creado correctamente", "Cuidador": CaregiverRequest})
}

func (h *CreateCaregiverController) ProcessCaregiver(Caregiver domain.Caregiver) error {
	fmt.Printf("Procesando creación de cuidador: %+v\n", Caregiver)
	return h.CreateCaregiver.Execute(Caregiver)
}
