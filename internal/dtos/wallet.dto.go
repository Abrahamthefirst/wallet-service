package dtos

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type TransferBetweenUsersRequestBody struct {
	ReceiverWalletID uint           `json:"receiver_wallet_id"`
	SenderWalletID   uint           `json:"sender_wallet_id"`
	IdempotencyKey   string         `json:"idempotency_key"`
	Description      string         `json:"description"`
	Amount           uint           `json:"amount"`
	Currency         enums.Currency `json:"currency"`
}

type CreateWalletRequestBody struct {
	Balance    uint             `json:"initial_balance"`
	Currency   enums.Currency   `json:"currency"`
	WalletType enums.WalletType `json:"wallet_type"`
}
type CreateWalletResponseBody struct {
	Message    string          `json:"message"`
	Data       entities.Wallet `json:"data"`
	StatusCode uint            `json:"statusCode"`
}
