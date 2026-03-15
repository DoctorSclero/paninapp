package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"pietroballarin.com/paninup-backend/internal/auth"
	"pietroballarin.com/paninup-backend/internal/types"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

func RequireRole(role types.UserRole) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists || userRole != role {
			ctx.JSON(403, gin.H{"error": "Forbidden: insufficient permissions"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
