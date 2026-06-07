package repository

import (
	"context"

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
func (r *WalletRepository) GetByID(ctx context.Context) {

}

func (r *WalletRepository) Update(ctx context.Context) {

}

func (r *WalletRepository) Create(ctx context.Context) {
}


func (r *WalletRepository) GetAll(ctx context.Context) {}
