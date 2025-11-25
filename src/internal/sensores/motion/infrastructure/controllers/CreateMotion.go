package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pulse_sense/src/core"
	patients "pulse_sense/src/internal/sensores/patients/domain"
	"pulse_sense/src/internal/sensores/motion/application"
	"pulse_sense/src/internal/sensores/motion/domain"
	fcm "pulse_sense/src/internal/services/fcm"
	websocket "pulse_sense/src/internal/services/websocket/application"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateMotionController struct {
	createMotion *application.CreateMotion
	wsService    *websocket.WebSocketService
	patientRepo  patients.PatientRepository
	userRepo     *core.UserRepository
	fcmSender    *fcm.FCMSender
}

func NewCreateMotionController(
	createMotion *application.CreateMotion,
	wsService *websocket.WebSocketService,
	patientRepo patients.PatientRepository,
	userRepo *core.UserRepository,
	fcmSender *fcm.FCMSender,
) *CreateMotionController {
	return &CreateMotionController{
		createMotion: createMotion,
		wsService:    wsService,
		patientRepo:     patientRepo,
		userRepo:     userRepo,
		fcmSender:    fcmSender,
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

	cage, err := h.patientRepo.GetPatientByID(fmt.Sprintf("%d", motion.IDPaciente))
	if err != nil {
		log.Printf("[ERROR] No se pudo obtener jaula %d: %v", motion.IDPaciente, err)
		return err
	}

	wsData := gin.H{
		"cage_id": motion.IDPaciente,
		"data": gin.H{
			"idpaciente":    motion.IDPaciente,
			"movimiento":    motion.Movimiento,
			"hora_registro": motion.HoraRegistro,
		},
		"event":     "new_motion",
		"timestamp": time.Now().Unix(),
	}

	if err := h.wsService.NotifyUser(cage.IDDoctor, wsData); err != nil {
		log.Printf("[WARN] Error notificando usuario %d via WebSocket: %v", cage.IDDoctor, err)
	}

	user, err := h.userRepo.GetUserByID(cage.IDDoctor)
	if err != nil {
		log.Printf("[ERROR] No se pudo obtener usuario %d: %v", cage.IDDoctor, err)
		return fmt.Errorf("error obteniendo usuario: %v", err)
	}

	if user.FCMToken != "" {
		status := "sin movimiento"
		if motion.Movimiento {
			status = "Movimiento detectado"
		}

		payload := fcm.NotificationPayload{
			Title: "Detección de movimiento",
			Body:  fmt.Sprintf("Jaula %d: %s", motion.IDPaciente, status),
			Data: map[string]string{
				"patient_id":    fmt.Sprintf("%d", motion.IDPaciente),
				"movimiento": fmt.Sprintf("%t", motion.Movimiento),
				"timestamp":  time.Now().Format(time.RFC3339),
				"event":      "new_motion",
			},
		}

		if err := h.fcmSender.SendNotification(context.Background(), user.FCMToken, payload); err != nil {
			log.Printf("[ERROR] Fallo al enviar notificación FCM: %v", err)
		}
	}

	log.Printf("[INFO] Notificación de movimiento enviada: %+v", wsData)
	return nil
}
