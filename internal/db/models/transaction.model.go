package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

// I need to come back here and change the transaction type to an enum it to an enum
type TransactionModel struct {
	gorm.Model      `gorm:"uniqueIndex"`
	IdempotencyKey  string
	WalletId        uint
	Amount          uint
	Currency        enums.Currency
	FinalBalance    uint
	Desscription    string
	TransactionType string
	Fee             FeeModel
}

func (*TransactionModel) TableName() string {
	return "transactions"
}
