package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/workers/application"

	"github.com/gin-gonic/gin"
)

type WorkerAssignmentController struct {
    CheckAssignmentService *application.CheckWorkerAssignment
}

func NewWorkerAssignmentController(checkAssignment *application.CheckWorkerAssignment) *WorkerAssignmentController {
    return &WorkerAssignmentController{
        CheckAssignmentService: checkAssignment,
    }
}

func (h *WorkerAssignmentController) CheckAssignment(c *gin.Context) {
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
            "message": "El trabajador ya está asignado a un hospital",
            "assigned": true,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "message": "El trabajador no está asignado a ningún hospital",
            "assigned": false,
        })
    }
}