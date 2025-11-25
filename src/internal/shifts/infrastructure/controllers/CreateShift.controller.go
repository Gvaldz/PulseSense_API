package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/shifts/application"
	"pulse_sense/src/internal/shifts/domain"

	"github.com/gin-gonic/gin"
)

type CreateShiftController struct {
	CreateShift *application.CreateShift
}

func NewCreateShiftController(CreateShift *application.CreateShift) *CreateShiftController {
	return &CreateShiftController{
		CreateShift: CreateShift,
	}
}

func (h *CreateShiftController) Create(c *gin.Context) {
	var ShiftRequest domain.Shift
	if err := c.ShouldBindJSON(&ShiftRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CreateShift.Execute(ShiftRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Turno creado correctamente", "Shift": ShiftRequest})
}

func (h *CreateShiftController) ProcessShift(Shift domain.Shift) error {
	fmt.Printf("Procesando creación de turno: %+v\n", Shift)
	return h.CreateShift.Execute(Shift)
}
