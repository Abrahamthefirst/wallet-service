package endpoints

import (
	"log/slog"
	"net/http"

	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
	custom_http "github.com/Abrahamthefirst/finecore-practice/internal/http/webutil"
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

	var requestDto dtos.SignupRequestDto

	if err := custom_http.ValidateRequest(c, &requestDto); err != nil {
		return
	}

	user, err := a.authService.SignUp(c, requestDto)

	if err != nil {
		custom_http.HandleError(c, err)
		return
	}

	response := dtos.SignupResponse{
		Message:    "Registration Successful",
		Data:       dtos.SignupResponseBody{User: *user},
		StatusCode: 201,
	}

	c.JSON(http.StatusOK, response)
}

func (a *AuthController) Login(c *gin.Context) {

	var requestDto dtos.LoginRequestDto

	if err := custom_http.ValidateRequest(c, &requestDto); err != nil {
		return

	}

	result, err := a.authService.Login(c.Request.Context(), requestDto)

	if err != nil {
		custom_http.HandleError(c, err)
		return
	}

	response := dtos.LoginOkResponse{
		Message:    "Login Successfull",
		Data:       dtos.LoginResponseBody{User: *result.User, AccessToken: result.AccessToken},
		StatusCode: 200,
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("refresh", result.RefreshToken, 3600*24, "/", "", true, true)
	c.JSON(http.StatusOK, response)

}

func RegisterAuthRoutes(r *gin.RouterGroup, authController *AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", authController.Signup)
		auth.POST("/login", authController.Login)
	}
}
