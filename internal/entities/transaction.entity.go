package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Transaction struct {
	ID             uint                `json:"id"`
	Amount         uint                `json:"amount"`
	Currency       enums.Currency      `json:"currency"`
	FinalBalance   uint                `json:"final_balance"`
	OperationType  enums.OperationType `json:"operation_type"`
	Description    string              `json:"description"`
	Fee            *Fee
	IdempotencyKey string `json:"idempotency_key"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
