package repository

import "gorm.io/gorm"

type LedgerRepository struct {
	db *gorm.DB
}

func NewLedgerRepository(db *gorm.DB) *LedgerRepository {

	return &LedgerRepository{
		db,
	}

}

func (r LedgerRepository) Create() {}
func (r LedgerRepository) GetByMerchantID() {}
func (r LedgerRepository) GetByID() {}
