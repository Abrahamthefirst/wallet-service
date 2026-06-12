package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

type LedgerEntryModel struct {
	gorm.Model
	TransactionID uint            `gorm:"not null;index"`
	AccountID     uint            `gorm:"not null;index"`
	EntryType     enums.EntryType `gorm:"not null"`
	Amount        uint             `gorm:"not null"`
	Currency      enums.Currency         `gorm:"not null;size:3"` // ISO 4217 e.g. "NGN"
	Description   string
}

func (*LedgerEntryModel) TableName() string {
	return "ledger_entries"
}
func (m *LedgerEntryModel) ToDomain() *entities.Ledger {
	ledgerEntry := &entities.Ledger{
		ID:            m.ID,
		TransactionID: m.TransactionID,
		Amount:        m.Amount,
		Currency:      m.Currency,
		EntryType:     m.EntryType,
		Description:   m.Description,
		UpdatedAt: m.UpdatedAt,
		CreatedAt: m.CreatedAt,
	}

	return ledgerEntry
}
