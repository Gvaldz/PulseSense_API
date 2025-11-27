package controllers

import (
	"fmt"
	"log"
	"net/http"
	"pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/motion/application"
	"pulse_sense/src/internal/sensores/motion/domain"
	patients "pulse_sense/src/internal/sensores/patients/domain"
	websocket "pulse_sense/src/internal/services/websocket/application"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateMotionController struct {
	createMotion *application.CreateMotion
	wsService    *websocket.WebSocketService
	patientRepo  patients.PatientRepository
	userRepo     *core.UserRepository
}

func NewCreateMotionController(
	createMotion *application.CreateMotion,
	wsService *websocket.WebSocketService,
	patientRepo patients.PatientRepository,
	userRepo *core.UserRepository,
) *CreateMotionController {
	return &CreateMotionController{
		createMotion: createMotion,
		wsService:    wsService,
		patientRepo:  patientRepo,
		userRepo:     userRepo,
	}
}

func (h *CreateMotionController) Create(c *gin.Context) {
	var motionRequest domain.Motion
	if err := c.ShouldBindJSON(&motionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Creando movimiento: %+v\n", motionRequest)

	err := h.createMotion.Execute(motionRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movimiento creado correctamente", "motion": motionRequest})
}

func (h *CreateMotionController) ProcessMotion(motion domain.Motion) error {
	log.Printf("[DEBUG] Iniciando procesamiento de movimiento: %+v", motion)

	if err := h.createMotion.Execute(motion); err != nil {
		log.Printf("[ERROR] Fallo al guardar movimiento: %v", err)
		return err
	}

	patient, err := h.patientRepo.GetPatientByID(fmt.Sprintf("%d", motion.IDPaciente))
	if err != nil {
		log.Printf("[ERROR] No se pudo obtener jaula %d: %v", motion.IDPaciente, err)
		return err
	}

	wsData := gin.H{
		"patient_id": motion.IDPaciente,
		"data": gin.H{
			"idpaciente":    motion.IDPaciente,
			"movimiento":    motion.Movimiento,
			"hora_registro": motion.HoraRegistro,
		},
		"event":     "new_motion",
		"timestamp": time.Now().Unix(),
	}

	if err := h.wsService.NotifyUser(patient.IDDoctor, wsData); err != nil {
		log.Printf("[WARN] Error notificando usuario %d via WebSocket: %v", patient.IDDoctor, err)
	}

	log.Printf("[INFO] Notificación de movimiento enviada: %+v", wsData)
	return nil
}
