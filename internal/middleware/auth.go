package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "makerble-assessment/internal/service"
)

func AuthMiddleware(authService *service.AuthService, requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        token := authHeader
        parts := strings.Split(authHeader, " ")
        if len(parts) == 2 && parts[0] == "Bearer" {
            token = parts[1]
        } else if len(parts) != 1 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }

        claims, err := authService.ValidateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        role, ok := claims["role"].(string)
        if !ok || role != requiredRole {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }

        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}