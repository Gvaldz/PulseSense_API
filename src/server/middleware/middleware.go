package middleware

import (
	"net/http"
	auth "pulse_sense/src/core"
	tokenService "pulse_sense/src/internal/services/auth/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService tokenService.TokenService, authRepo *auth.AuthRepository, allowedTypes ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token de autorización requerido"})
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "formato de token inválido"})
            return
        }

        userID, userType, err := tokenService.ValidateToken(parts[1])
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
            return
        }

        if len(allowedTypes) > 0 {
            typeAllowed := false
            for _, allowedType := range allowedTypes {
                if userType == allowedType {
                    typeAllowed = true
                    break
                }
            }
            if !typeAllowed {
                c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "acceso no autorizado para este tipo de usuario"})
                return
            }
        }

        user, err := authRepo.FindUserByID(userID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar usuario"})
            return
        }

        c.Set("userID", userID)
        c.Set("userType", userType)
        c.Set("user", user)
        c.Next()
    }
}