package repository

import (
	"context"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/models"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db,
	}
}

// GetByID retrieves a wallet by ID
func (r *WalletRepository) GetByID(ctx context.Context, id uint) (*entities.Wallet, error) {

	var wallet models.WalletModel

	err := DBFromCtx(ctx, r.db).Where("id = ?", id).First(&wallet).Error

	if err != nil {
		return nil, err
	}
	return wallet.ToDomain(), err

}

func (r *WalletRepository) Update(ctx context.Context) {

}

func (r *WalletRepository) Create(ctx context.Context) {
}

func (r *WalletRepository) GetAll(ctx context.Context) {}
