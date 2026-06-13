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
		AccountID:     input.AccountID,
		EntryType:     input.EntryType,
		Amount:        input.Amount,
		Currency:      input.Currency,
		Description:   input.Description,
	}
	err := DBFromCtx(ctx, r.db).Create(&ledgerEntry).Error
	if err != nil {
		return nil, err
	}
	return ledgerEntry.ToDomain(), nil
}

func (r LedgerRepository) CreateInBatch(ctx context.Context, input []entities.Ledger) ([]entities.Ledger, error) {

	ledgerEntriesList := make([]models.LedgerEntryModel, 0, len(input))

	for _, ledgerEntry := range input {
		ledgerEntryModel := models.LedgerEntryModel{
			TransactionID: ledgerEntry.TransactionID,
			AccountID:     ledgerEntry.AccountID,
			EntryType:     ledgerEntry.EntryType,
			Amount:        ledgerEntry.Amount,
			Currency:      ledgerEntry.Currency,
			Description:   ledgerEntry.Description,
		}
		ledgerEntriesList = append(ledgerEntriesList, ledgerEntryModel)
	}

	err := DBFromCtx(ctx, r.db).CreateInBatches(&ledgerEntriesList, 100).Error
	if err != nil {
		return nil, err
	}

	domainLedgers := make([]entities.Ledger, 0, len(ledgerEntriesList))
	for _, model := range ledgerEntriesList {
		domainLedgers = append(domainLedgers, *model.ToDomain())
	}

	return domainLedgers, nil
}
func (r LedgerRepository) GetByMerchantID() {}
func (r LedgerRepository) GetByID()         {}
