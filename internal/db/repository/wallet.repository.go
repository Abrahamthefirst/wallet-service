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

func (r *WalletRepository) UpdateBalance(ctx context.Context, walletId, newBalance int) (*entities.Wallet, error) {
	var wallet models.WalletModel

	err := DBFromCtx(ctx, r.db).Where("id = ?", walletId).Update("balance", newBalance).Error

	if err != nil {
		return nil, err
	}
	return wallet.ToDomain(), err

}

func (r *WalletRepository) Create(ctx context.Context, userId uint, input entities.Wallet) (*entities.Wallet, error) {
	wallet := models.WalletModel{
		UserId:     input.UserId,
		Balance:    input.Balance,
		Currency:   input.Currency,
		WalletType: input.WalletType,
	}
	err := DBFromCtx(ctx, r.db).Create(&wallet).Error
	if err != nil {
		return nil, err
	}
	return wallet.ToDomain(), nil
}

func (r *WalletRepository) GetAll(ctx context.Context) (*[]entities.Wallet, error) {
	var wallets []models.WalletModel

	err := DBFromCtx(ctx, r.db).Find(&wallets).Error

	if err != nil {
		return nil, err
	}

	walletsList := []entities.Wallet{}

	for _, wallet := range wallets {
		walletsList = append(walletsList, *wallet.ToDomain())
	}

	return &walletsList, nil
}

func (r *WalletRepository) FetchByUserID(ctx context.Context, userID uint) (*[]entities.Wallet, error) {
	var wallets []models.WalletModel

	// We are querying the wallets table directly
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&wallets).Error
	if err != nil {
		return nil, err
	}

	walletsList := []entities.Wallet{}

	for _, wallet := range wallets {
		walletsList = append(walletsList, *wallet.ToDomain())
	}
	return &walletsList, nil
}
