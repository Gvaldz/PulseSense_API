package controllers

import (
	"net/http"
	"pulse_sense/src/internal/hospitals/application"

	"github.com/gin-gonic/gin"
)

type SearchHospitalController struct {
	searchHospital *application.SearchHospital
}

func NewSearchHospitalController(searchHospital *application.SearchHospital) *SearchHospitalController {
	return &SearchHospitalController{searchHospital: searchHospital}
}

func (h *SearchHospitalController) SearchHospital(c *gin.Context) {
	name := c.Param("name")

	user, err := h.searchHospital.Execute(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
