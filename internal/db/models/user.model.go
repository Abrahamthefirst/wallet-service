package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model   `gorm:"uniqueIndex"`
	FirstName    string
	LastName     string
	UserName     string
	PasswordHash string
	ProfileUrl   string
	wallets      []WalletModel
}

func (*UserModel) TableName() string {
	return "users"
}
