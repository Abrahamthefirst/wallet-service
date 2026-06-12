package dtos

import (
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
