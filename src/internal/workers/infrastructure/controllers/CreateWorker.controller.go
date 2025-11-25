package controllers

import (
	"fmt"
	"net/http"
	"pulse_sense/src/internal/workers/application"
	"pulse_sense/src/internal/workers/domain"

	"github.com/gin-gonic/gin"
)

type CreateWorkerController struct {
	CreateWorker *application.CreateWorker
}

func NewCreateWorkerController(CreateWorker *application.CreateWorker) *CreateWorkerController {
	return &CreateWorkerController{
		CreateWorker: CreateWorker,
	}
}

func (h *CreateWorkerController) Create(c *gin.Context) {
	var WorkerRequest domain.Worker
	if err := c.ShouldBindJSON(&WorkerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.CreateWorker.Execute(WorkerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Worker creado correctamente", "Worker": WorkerRequest})
}

func (h *CreateWorkerController) ProcessWorker(Worker domain.Worker) error {
	fmt.Printf("Procesando creación de Worker: %+v\n", Worker)
	return h.CreateWorker.Execute(Worker)
}
