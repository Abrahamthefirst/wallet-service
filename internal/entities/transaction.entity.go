package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Transaction struct {
	ID              uint `json:"id"`
	WalletId        uint
	Amount          uint
	Currency        enums.Currency
	FinalBalance    uint
	Desscription    string
	TransactionType string

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
