package endpoints

import (
	"log/slog"

	authservice "github.com/Abrahamthefirst/finecore-practice/internal/service/auth-service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger      *slog.Logger
	authService *authservice.AuthService
}

func NewAuthController(authSerivce *authservice.AuthService) *AuthController {
	return &AuthController{
		logger:      slog.Default().With("component", "auth"),
		authService: authSerivce,
	}
}

func (a *AuthController) Signup(c *gin.Context) {

}

func RegisterAuthorRoutes(r *gin.Engine, authController *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("", authController.Signup)
	}
}
