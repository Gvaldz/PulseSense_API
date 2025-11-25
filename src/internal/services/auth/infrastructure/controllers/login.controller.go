package controllers

import (
	"net/http"
	"pulse_sense/src/internal/services/auth/application"
	"pulse_sense/src/internal/users/domain"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUC *application.Login
}

func NewLoginController(loginUC *application.Login) *LoginController {
	return &LoginController{loginUC: loginUC}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var credentials domain.User
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "petición inválida"})
		return
	}

	token, userType, err := c.loginUC.Execute(credentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+token.Token)
	ctx.JSON(http.StatusOK, gin.H{
		"token":      token.Token,
		"expires_at": token.ExpiresAt,
		"user_type":  userType,
	})
}
