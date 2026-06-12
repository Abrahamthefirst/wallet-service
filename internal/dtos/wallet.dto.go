package dtos

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type TransferBetweenUsersRequestBody struct {
	ReceiverWalletID uint
	SenderWalletID   uint
	IdempotencyKey   string
	Description      string
	Amount           uint
	Currency         enums.Currency
}
