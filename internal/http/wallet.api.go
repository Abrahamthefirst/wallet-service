package endpoints

import (
	"log/slog"

	walletservice "github.com/Abrahamthefirst/finecore-practice/internal/service/wallet-service"
	"github.com/gin-gonic/gin"
)

type WalletController struct {
	logger        *slog.Logger
	walletService *walletservice.WalletService
}

func NewWalletController(walletService *walletservice.WalletService) *WalletController {
	return &WalletController{
		logger:        slog.Default().With("component", "wallet"),
		walletService: walletService,
	}
}

func (w *WalletController) Fund(c *gin.Context) {

}

func RegisterWalletRoutes(r *gin.Engine, walletController *WalletController) {
	wallet := r.Group("/wallets")
	{
		wallet.POST("/:id/fund", walletController.Fund)
		// auth.GET("/:id", walletController.GetAuthorByID)
	}
}
