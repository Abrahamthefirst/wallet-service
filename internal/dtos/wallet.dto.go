package dtos

import "github.com/Abrahamthefirst/finecore-practice/internal/enums"

type TransferBetweenUsersRequestBody struct {
	ReceiverId uint
	Amount     uint
	Currency   enums.Currency
}
