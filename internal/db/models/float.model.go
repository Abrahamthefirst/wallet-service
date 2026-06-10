package models

import (
	"github.com/Abrahamthefirst/finecore-practice/internal/enums"
	"gorm.io/gorm"
)

type FloatModel struct {
	gorm.Model    `gorm:"uniqueIndex"`
	TransactionId uint
	Amount        uint
	Percentage    uint
	Currency      enums.Currency
	WalletType    enums.WalletType
}

func (*FloatModel) TableName() string {
	return "platform_fee"
}
