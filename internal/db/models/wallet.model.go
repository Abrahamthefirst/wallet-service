package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

type WalletModel struct {
	gorm.Model   `gorm:"uniqueIndex"`
	UserId       uint
	Balance      uint
	Currency     enums.Currency
	WalletType   enums.WalletType
	transactions []TransactionModel
}

func (*WalletModel) TableName() string {
	return "wallets"
}

func (m *WalletModel) ToDomain() *entities.Wallet {
	wallet := &entities.Wallet{
		ID:         m.ID,
		Balance:    m.Balance,
		UserId:     m.UserId,
		Currency:   m.Currency,
		WalletType: m.WalletType,
		UpdatedAt:  m.UpdatedAt,
		CreatedAt:  m.CreatedAt,
	}

	if wallet.Transactions == nil {
		wallet.Transactions = &[]entities.Transaction{}
	}

	iterationStep := 0
	if len(m.transactions) > 0 {

		for _, item := range m.transactions {

			transaction := entities.Transaction{
				ID:              item.ID,
				WalletId:        item.WalletId,
				Amount:          item.Amount,
				Currency:        item.Currency,
				FinalBalance:    item.FinalBalance,
				Description:    item.Description,
				OperationType: item.OperationType,
				CreatedAt:       item.CreatedAt,
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
