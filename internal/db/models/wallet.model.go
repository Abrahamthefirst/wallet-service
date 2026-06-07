package models

import (
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
