package controllers

import (
	"fmt"
	"log"
	"net/http"
	core "pulse_sense/src/core"
	"pulse_sense/src/internal/sensores/signos/application"
	"pulse_sense/src/internal/sensores/signos/domain"
	websocket "pulse_sense/src/internal/services/websocket/application"
	patient "pulse_sense/src/internal/sensores/patients/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateSignsController struct {
	createSign  *application.CreateSigns
	wsService   *websocket.WebSocketService
	patientRepo patient.PatientRepository
	userRepo    *core.UserRepository
}

func NewCreateSignsController(
	createSign *application.CreateSigns,
	wsService *websocket.WebSocketService,
	patientRepo patient.PatientRepository,
	userRepo *core.UserRepository,
) *CreateSignsController {
	return &CreateSignsController{
		createSign:  createSign,
		wsService:   wsService,
		patientRepo: patientRepo,
		userRepo:    userRepo,
	}
}

func (h *CreateSignsController) Create(c *gin.Context) {
	var SignRequest domain.Sign
	if err := c.ShouldBindJSON(&SignRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Creando signos desde HTTP: %+v\n", SignRequest)

	err := h.createSign.Execute(SignRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "signos creada correctamente", "Sign": SignRequest})
}

func (h *CreateSignsController) ProcessSign(Sign domain.Sign) error {
	log.Printf("[DEBUG] Iniciando procesamiento de signos: %+v", Sign)

	if err := h.createSign.Execute(Sign); err != nil {
		log.Printf("[ERROR] Fallo al guardar signos: %v", err)
		return err
	}
	log.Printf("[DEBUG] signos guardada en BD: %+v", Sign)

	patient, err := h.patientRepo.GetPatientByID(fmt.Sprintf("%d", Sign.IDPaciente))
	if err != nil {
		log.Printf("[ERROR] No se pudo obtener Paciente %d: %v", Sign.IDPaciente, err)
		return err
	}
	log.Printf("[DEBUG] Paciente obtenida: %+v", patient)

	wsData := gin.H{
		"event":      "new_Sign",
		"data":       Sign,
		"patient_id": fmt.Sprintf("%d", Sign.IDPaciente),
		"timestamp":  time.Now().Unix(),
	}

	if err := h.wsService.NotifyUser(patient.IDDoctor, wsData); err != nil {
		log.Printf("[WARN] Error notificando usuario %d via WebSocket: %v", patient.IDDoctor, err)
	} else {
		log.Printf("[DEBUG] Notificación WebSocket enviada al usuario %d", patient.IDDoctor)
	}


	log.Printf("[INFO] Procesamiento completado para signos: %+v", Sign)
	return nil
}
