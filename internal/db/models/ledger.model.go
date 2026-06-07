package models

import "gorm.io/gorm"

type LedgerEntryModel struct {
	gorm.Model `gorm:"uniqueIndex"`
	TransactionId uint
	WalletId uint
	UserId uint
	EntryType string
	Amount uint
	Currency string
	Description string
}


func (*LedgerEntryModel) TableName() string {
	return "ledger_entries"
}