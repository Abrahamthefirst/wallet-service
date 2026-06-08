package models

import (
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model      `gorm:"uniqueIndex"`
	FirstName       *string
	LastName        *string
	Username        string
	Email           string
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	Password        string
	AvatarKey       *string
	wallets         []WalletModel
}

func (*UserModel) TableName() string {
	return "users"
}

func (m *UserModel) ToDomain() *entities.User {
	return &entities.User{
		ID:              m.ID,
		Email:           m.Email,
		Username:        m.Username,
		AvatarKey:       *m.AvatarKey,
		Password:        m.Password,
		EmailVerifiedAt: m.EmailVerifiedAt,
		CreatedAt:       m.CreatedAt,
	}
}
