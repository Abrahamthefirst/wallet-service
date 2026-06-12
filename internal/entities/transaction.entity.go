package entities

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
)

type Transaction struct {
	ID             uint `json:"id"`
	Amount         uint
	Currency       enums.Currency
	FinalBalance   uint
	OperationType  enums.OperationType
	Description    string
	Fee            *Fee
	IdempotencyKey string

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
