package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pulse_sense/src/internal/userpatient/application"
)

type GetNurseByPatientController struct {
	usecase *application.GetNursesByPatientId
}

func NewGetNurseByPatientController(usecase *application.GetNursesByPatientId) *GetNurseByPatientController {
	return &GetNurseByPatientController{usecase: usecase}
}

func (c *GetNurseByPatientController) GetByCuidador(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	results, err := c.usecase.Execute(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}
