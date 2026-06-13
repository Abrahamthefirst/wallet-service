package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

type WalletModel struct {
	gorm.Model
	OwnerId      uint               `gorm:"not null;index"`
	Balance      uint               `gorm:"not null;default:0"`
	Currency     enums.Currency     `gorm:"not null"`
	WalletType   enums.WalletType   `gorm:"not null"`
	Transactions []TransactionModel `gorm:"foreignKey:WalletId"`
}

func (*WalletModel) TableName() string {
	return "wallets"
}

func (m *WalletModel) ToDomain() *entities.Wallet {
	wallet := &entities.Wallet{
		ID:         m.ID,
		Balance:    m.Balance,
		UserId:     m.OwnerId,
		Currency:   m.Currency,
		WalletType: m.WalletType,
		UpdatedAt:  m.UpdatedAt,
		CreatedAt:  m.CreatedAt,
	}

	if wallet.Transactions == nil {
		wallet.Transactions = &[]entities.Transaction{}
	}

	iterationStep := 0
	if len(m.Transactions) > 0 {

		for _, item := range m.Transactions {

			transaction := entities.Transaction{
				ID:            item.ID,
				Amount:        item.Amount,
				Currency:      item.Currency,
				FinalBalance:  item.FinalBalance,
				Description:   item.Description,
				OperationType: item.OperationType,
				CreatedAt:     item.CreatedAt,
			}

			*wallet.Transactions = append(*wallet.Transactions, transaction)

			iterationStep++

			if iterationStep >= 10 {
				break
			}
		}
	}
	return wallet
}
