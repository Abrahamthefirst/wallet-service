package middleware

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        header := c.GetHeader("Authorization")

        if header == "" || !strings.HasPrefix(header, "Bearer ") {
            c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
            return
        }

        tokenString := strings.TrimPrefix(header, "Bearer ")
        
        // 1. Declare your claims struct variable here
        claims := &entities.TokenClaims{}

        // 2. Pass 'claims' directly into the function
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
        })

        if err != nil {
            slog.Error("Parse error: " + err.Error())
            c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
            return
        }

        // 3. Since 'claims' was passed by reference, jwt.ParseWithClaims populated it.
        // We just check token.Valid without complex type assertions.
        if token.Valid {
            slog.Info("User authenticated", "id", claims.UserID)
            c.Set("user_id", claims.UserID)
            c.Next()
            return
        }

        // Catch-all safety fallback
        c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
    }
}