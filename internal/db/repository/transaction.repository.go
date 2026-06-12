package repository

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/models"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
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

func (r *TransactionRepository) GetByID(ctx context.Context, id uint) (*entities.Transaction, error) {
	var transaction models.TransactionModel
	err := DBFromCtx(ctx, r.db).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction.ToDomain(), nil

}

func (r *TransactionRepository) Create(ctx context.Context, input entities.Transaction) (*entities.Transaction, error) {
	transaction := models.TransactionModel{
		FinalBalance:   input.FinalBalance,
		Currency:       input.Currency,
		Description:    input.Description,
		IdempotencyKey: input.IdempotencyKey,
	}
	err := DBFromCtx(ctx, r.db).Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction.ToDomain(), nil
}
