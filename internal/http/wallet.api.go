package endpoints

import (
	"log/slog"
	"net/http"

	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
	custom_http "github.com/Abrahamthefirst/finecore-practice/internal/http/webutil"
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

func (w *WalletController) CreateWallet(c *gin.Context) {

	var requestDto dtos.CreateWalletRequestBody

	if err := custom_http.ValidateRequest(c, &requestDto); err != nil {
		return
	}

	userID, ok := c.Get("user_id")
	userId, ok := userID.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, custom_http.ErrorResponse{
			Message:    "Not Authenticated",
			StatusCode: http.StatusUnauthorized,
		})
	}
	wallet, err := w.walletService.CreateWallet(c, userId, &requestDto)

	if err != nil {
		custom_http.HandleError(c, err)
		return
	}

	response := dtos.CreateWalletResponseBody{
		Message:    "Registration Successful",
		Data:       *wallet,
		StatusCode: 201,
	}

	c.JSON(http.StatusOK, response)
}

func (w *WalletController) GetAllWallets(c *gin.Context) {

}

func (w *WalletController) GetWalletById(c *gin.Context) {

}

func (w *WalletController) GetWalletTransactions(c *gin.Context) {

}
func RegisterWalletRoutes(r *gin.RouterGroup, walletController *WalletController) {
	wallet := r.Group("/wallets")
	{
		wallet.POST("", walletController.CreateWallet)
		wallet.GET("", walletController.GetAllWallets)
		wallet.GET(":id", walletController.GetWalletById)

		// transfers, deposits and withdrawals should have their own api
		// auth.GET("/:id", walletController.GetAuthorByID)
	}
}
