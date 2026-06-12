package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

type LedgerEntryModel struct {
	gorm.Model `gorm:"uniqueIndex"`
	TransactionId uint
	WalletId uint
	UserId uint
	EntryType enums.EntryType
	Amount uint
	Currency string
	Description string
}


func (*LedgerEntryModel) TableName() string {
	return "ledger_entries"
}