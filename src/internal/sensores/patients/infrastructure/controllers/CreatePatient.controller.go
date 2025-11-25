package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/sensores/patients/application"
	"pulse_sense/src/internal/sensores/patients/domain"

	"github.com/gin-gonic/gin"
)

type CreatePatientController struct {
	CreatePatient *application.CreatePatient
}

func NewCreatePatientController(CreatePatient *application.CreatePatient) *CreatePatientController {
	return &CreatePatientController{
		CreatePatient: CreatePatient,
	}
}

func (h *CreatePatientController) Create(c *gin.Context) {
    var patientRequest domain.Patient
    if err := c.ShouldBindJSON(&patientRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := h.CreatePatient.Execute(patientRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    patientRequest.IdPaciente = int32(id)
    
    c.JSON(http.StatusCreated, gin.H{
        "message": "Paciente creado correctamente", 
        "patient": patientRequest,
        "id": id,
    })
}
func (h *CreatePatientController) Processpatient(patient domain.Patient) error {
    fmt.Printf("Procesando creación de Paciente: %+v\n", patient)
    _, err := h.CreatePatient.Execute(patient)
    return err
}
