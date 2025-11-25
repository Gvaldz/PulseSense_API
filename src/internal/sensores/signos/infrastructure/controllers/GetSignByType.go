package controllers

import (
	"net/http"
	"pulse_sense/src/internal/sensores/signos/application"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetSignsByTypeAndTimeController struct {
    getByTypeAndTime *application.GetSignsByTypeAndTimeRange
}

func NewGetSignsByTypeAndTimeController(getByTypeAndTime *application.GetSignsByTypeAndTimeRange) *GetSignsByTypeAndTimeController {
    return &GetSignsByTypeAndTimeController{getByTypeAndTime: getByTypeAndTime}
}

func (h *GetSignsByTypeAndTimeController) GetByTypeAndTime(c *gin.Context) {
    IDPacienteStr := c.Param("IDPaciente")
    IDTipoStr := c.Param("IDTipo") 
    fecha := c.Param("fecha")
    turno := strings.ToLower(c.Param("turno")) 
    
    validTurnos := map[string]bool{"matutino": true, "vespertino": true, "nocturno": true}
    if !validTurnos[turno] {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Turno no válido. Use matutino, vespertino o nocturno"})
        return
    }
    
    IDTipo, err := strconv.Atoi(IDTipoStr)
    IDPaciente, err := strconv.Atoi(IDPacienteStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de paciente o tipo inválido"})
        return
    }
    
    signs, err := h.getByTypeAndTime.Execute(IDPaciente, IDTipo, fecha, turno)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, signs)
}