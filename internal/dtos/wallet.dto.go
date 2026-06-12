package dtos

import "github.com/Abrahamthefirst/finecore-practice/internal/enums"

type TransferBetweenUsersRequestBody struct {
	ReceiverWalletId uint
	SenderWalletId uint
	Amount     uint
	Currency   enums.Currency
}
