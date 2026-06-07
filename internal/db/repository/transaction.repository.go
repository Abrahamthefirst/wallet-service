package repository

import (
	"context"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db,
	}
}

func (r *TransactionRepository) GetByID(ctx context.Context) {

}

func (r *TransactionRepository) Create(ctx context.Context) {

}


