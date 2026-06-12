package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Ledger struct {
	ID            uint            `json:"id"`
	TransactionID uint            `json:"transaction_id"`
	AccountID     uint            `json:"account_id"`
	EntryType     enums.EntryType `json:"entry_type"`
	Amount        uint            `json:"amount"`
	Currency      enums.Currency  `json:"currency"`
	Description   string          `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
