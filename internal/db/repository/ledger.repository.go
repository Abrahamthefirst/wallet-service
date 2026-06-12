package repository

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/models"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"gorm.io/gorm"
)

type LedgerRepository struct {
	db *gorm.DB
}

func NewLedgerRepository(db *gorm.DB) *LedgerRepository {

	return &LedgerRepository{
		db,
	}

}

func (r LedgerRepository) Create(ctx context.Context, input entities.Ledger) (*entities.Ledger, error) {
	ledgerEntry := models.LedgerEntryModel{
		TransactionID: input.TransactionID,
		AccountID: input.AccountID,
		EntryType: input.EntryType,
		Amount:         input.Amount,
		Currency:       input.Currency,
		Description:    input.Description,

	}
	err := DBFromCtx(ctx, r.db).Create(&ledgerEntry).Error
	if err != nil {
		return nil, err
	}
	return ledgerEntry.ToDomain(), nil
}

func (r LedgerRepository) GetByMerchantID() {}
func (r LedgerRepository) GetByID()         {}
