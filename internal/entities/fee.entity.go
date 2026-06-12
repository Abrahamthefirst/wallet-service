package entities

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Fee struct {
	ID            uint             `json:"id"`
	TransactionId uint             `json:"transaction_id"`
	Amount        uint             `json:"amount"`
	Percentage    uint             `json:"percentage"`
	Currency      enums.Currency   `json:"currency"`
	WalletType    enums.WalletType `json:"wallet_type"`
}
