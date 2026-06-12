package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Ledger struct {
	ID            uint            `json:"id"`
	TransactionID uint          
	AccountID     uint           
	EntryType     enums.EntryType
	Amount        uint            
	Currency      enums.Currency 
	Description   string

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
