package entities

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Fee struct {
	ID            uint `json:"id"`
	TransactionId uint
	Amount        uint
	Percentage    uint
	Currency      enums.Currency
	WalletType    enums.WalletType
	Fee           uint
}
