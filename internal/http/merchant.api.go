package endpoints

import (
	"log/slog"

	merchantservice "github.com/Abrahamthefirst/finecore-practice/internal/service/merchant-service"
	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	logger          *slog.Logger
	merchantService *merchantservice.MerchantService
}

func NewMerchantController(merchantService *merchantservice.MerchantService) *MerchantController {
	return &MerchantController{
		logger:          slog.Default().With("component", "merchant"),
		merchantService: merchantService,
	}
}

func (m *MerchantController) InitiatePayout(c *gin.Context) {

}

func RegisterMerchantRoutes(r *gin.Engine, merchantController *MerchantController) {
	merchant := r.Group("merchant")
	{
		merchant.POST("/initiate-payout", merchantController.InitiatePayout)
		// auth.GET("/:id", walletController.GetAuthorByID)
	}
}
