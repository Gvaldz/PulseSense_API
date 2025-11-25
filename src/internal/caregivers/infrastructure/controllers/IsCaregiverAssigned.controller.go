package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/caregivers/application"

	"github.com/gin-gonic/gin"
)

type CaregiverAssignmentController struct {
    CheckAssignmentService *application.CheckCaregiverAssignment
}

func NewCaregiverAssignmentController(checkAssignment *application.CheckCaregiverAssignment) *CaregiverAssignmentController {
    return &CaregiverAssignmentController{
        CheckAssignmentService: checkAssignment,
    }
}

func (h *CaregiverAssignmentController) CheckAssignment(c *gin.Context) {
    idUsuario := c.Param("id") 
    
    var id int
    if _, err := fmt.Sscan(idUsuario, &id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    isAssigned, err := h.CheckAssignmentService.Execute(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if isAssigned {
        c.JSON(http.StatusOK, gin.H{
            "message": "El enfermero ya está asignado a un paciente",
            "assigned": true,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "message": "El enfermero no está asignado a ningún paciente",
            "assigned": false,
        })
    }
}