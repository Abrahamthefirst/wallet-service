package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

// I need to come back here and change the transaction type to an enum it to an enum
type TransactionModel struct {
	gorm.Model     `gorm:"uniqueIndex"`
	IdempotencyKey string
	Amount         uint `gorm:"not null"`
	Currency       enums.Currency
	FinalBalance   uint
	Description    string
	OperationType  enums.OperationType
}

func (*TransactionModel) TableName() string {
	return "transactions"
}

func (m *TransactionModel) ToDomain() *entities.Transaction {
	transaction := &entities.Transaction{
		ID:             m.ID,
		Amount:         m.Amount,
		Currency:       m.Currency,
		FinalBalance:   m.FinalBalance,
		OperationType:  m.OperationType,
		IdempotencyKey: m.IdempotencyKey,
		Description:    m.Description,

		UpdatedAt: m.UpdatedAt,
		CreatedAt: m.CreatedAt,
	}

	return transaction
}
